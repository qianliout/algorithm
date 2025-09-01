package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

// ProducePaymentSuccessTransaction sends four logically-related events in a single Kafka transaction.
// Topics:
//   - update order status
//   - deduct inventory
//   - send notification
//   - write audit log
//
// Parameters:
//   - brokers: Kafka broker list, e.g. []string{"localhost:9092"}
//   - transactionalID: globally unique transactional.id for this producer instance (stable across restarts of the same instance)
//   - key: message key to control partitioning if desired
//   - orderID: business identifier included in messages
//   - payloads: four payload strings corresponding to the four topics
//   - topics: four topic names in the order listed above
//
// Note: Consumer side should set isolation.level=read_committed to read only committed records.
func ProducePaymentSuccessTransaction(ctx context.Context, brokers []string, transactionalID string, key string, orderID string, payloads [4]string, topics [4]string) error {
	cfg := sarama.NewConfig()
	// Kafka 2.5+ recommended for transactions; adjust if your cluster is older.
	cfg.Version = sarama.V2_5_0_0

	// Exactly-once/transactions preconditions
	cfg.Producer.Idempotent = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	cfg.Net.MaxOpenRequests = 1

	// Transactional configuration
	cfg.Producer.Transaction = &sarama.TransactionConfig{ID: transactionalID}

	// Use a sync producer for simplicity; it's a thin wrapper over the async producer
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return fmt.Errorf("create producer: %w", err)
	}
	defer producer.Close()

	// Begin transaction
	if err := producer.BeginTxn(); err != nil {
		return fmt.Errorf("begin txn: %w", err)
	}

	// Helper to abort on any error
	abort := func(sendErr error) error {
		_ = producer.AbortTxn()
		return sendErr
	}

	// Build messages
	now := time.Now().UTC().Format(time.RFC3339Nano)
	msgs := []*sarama.ProducerMessage{
		{Topic: topics[0], Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(fmt.Sprintf("order=%s time=%s payload=%s", orderID, now, payloads[0]))},
		{Topic: topics[1], Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(fmt.Sprintf("order=%s time=%s payload=%s", orderID, now, payloads[1]))},
		{Topic: topics[2], Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(fmt.Sprintf("order=%s time=%s payload=%s", orderID, now, payloads[2]))},
		{Topic: topics[3], Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(fmt.Sprintf("order=%s time=%s payload=%s", orderID, now, payloads[3]))},
	}

	// Optional: simple context deadline for all sends
	deadline, hasDeadline := ctx.Deadline()

	for _, m := range msgs {
		if hasDeadline && time.Until(deadline) <= 0 {
			return abort(fmt.Errorf("context deadline exceeded before send"))
		}
		if _, _, err := producer.SendMessage(m); err != nil {
			return abort(fmt.Errorf("send to topic %s: %w", m.Topic, err))
		}
	}

	if err := producer.CommitTxn(); err != nil {
		return fmt.Errorf("commit txn: %w", err)
	}
	return nil
}
