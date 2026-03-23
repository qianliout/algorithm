package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// 统一接收退出信号，确保示例可被 Ctrl+C 优雅中断。
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runDemo(ctx); err != nil {
		log.Printf("transaction demo failed: %v", err)
		os.Exit(1)
	}
	log.Println("transaction demo finished successfully")
}

func runDemo(ctx context.Context) error {
	// 本地 Docker Kafka 映射端口为 9092。
	brokers := []string{"localhost:9092"}
	// 使用唯一 topic，避免历史数据干扰学习结果。
	topic := fmt.Sprintf("txn_demo_%d", time.Now().UnixNano())
	// transactional.id 必须稳定且唯一，这里用时间戳简化演示。
	transactionID := fmt.Sprintf("txn_producer_%d", time.Now().UnixNano())

	if err := createTopic(ctx, brokers, topic); err != nil {
		return err
	}

	committedMessageCh := make(chan string, 1)
	consumerErrCh := make(chan error, 1)
	go func() {
		// 消费端使用 read_committed，只会看到已提交事务消息。
		consumerErrCh <- consumeOneCommittedMessage(ctx, brokers, topic, committedMessageCh)
	}()

	// 给 consumer 一点启动时间，避免演示中错过最早消息。
	time.Sleep(500 * time.Millisecond)

	if err := produceTransaction(ctx, brokers, topic, transactionID); err != nil {
		return err
	}

	select {
	case msg := <-committedMessageCh:
		log.Printf("Consumer read committed message: %s", msg)
	case err := <-consumerErrCh:
		if err != nil {
			return err
		}
		return errors.New("consumer exited without committed message")
	case <-time.After(8 * time.Second):
		return errors.New("consumer timeout")
	}

	log.Println("Expected result: committed data is visible, aborted data is invisible")
	return nil
}

func createTopic(ctx context.Context, brokers []string, topic string) (err error) {
	adminConfig := sarama.NewConfig()
	adminConfig.Version = sarama.V3_6_0_0
	admin, err := sarama.NewClusterAdmin(brokers, adminConfig)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := admin.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	// 单机学习环境允许 topic 已存在，避免重复运行时报错。
	err = admin.CreateTopic(topic, topicDetail, false)
	if err != nil && !errors.Is(err, sarama.ErrTopicAlreadyExists) {
		return err
	}
	log.Printf("Topic ready: %s", topic)
	return nil
}

func produceTransaction(ctx context.Context, brokers []string, topic, transactionID string) (err error) {
	producerConfig := sarama.NewConfig()
	producerConfig.Version = sarama.V3_6_0_0
	// 必须：事务场景下建议使用 WaitForAll(-1)，只有所有 ISR 副本确认后才算成功，
	// 配合幂等与事务可以最大程度避免“已返回成功但副本未落盘”的风险。
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	// 必须（SyncProducer）：同步生产者内部依赖 Successes 通道拿到发送结果，
	// 若不为 true，SyncProducer 无法正常工作。
	producerConfig.Producer.Return.Successes = true
	// 必须（事务前提）：开启幂等生产，Kafka 用 PID+Sequence 去重，
	// 避免网络重试导致重复写入；Kafka 事务依赖幂等能力。
	producerConfig.Producer.Idempotent = true
	// 必须（与幂等强相关）：单连接上同一时刻只允许 1 个 in-flight 请求，
	// 避免重试后乱序导致 sequence 失配（尤其是旧版本 broker 更敏感）。
	producerConfig.Net.MaxOpenRequests = 1
	// 建议（非事务必需）：允许可恢复错误自动重试，提升成功率。
	// 是否重试不决定“是否是事务”，但会影响事务内发送稳定性。
	producerConfig.Producer.Retry.Max = 5
	// 必须（事务开关）：只要设置 Transaction.ID，这个 producer 才是事务生产者。
	// 同一 transactional.id 用于故障恢复与 fencing（隔离旧实例）。
	producerConfig.Producer.Transaction.ID = transactionID
	// 建议（事务运行参数）：单个事务允许的最长执行时间，超时后 broker 会中止事务。
	// 该值不是“开启事务”的必要条件，但会影响长事务行为。
	producerConfig.Producer.Transaction.Timeout = 20 * time.Second
	// sarama.NewAsyncProducer()
	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := producer.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// 事务 1：发送并提交，消息应当可见。
	if err = producer.BeginTxn(); err != nil {
		return err
	}

	committed := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("k1"),
		Value: sarama.StringEncoder("committed"),
	}
	if _, _, err = producer.SendMessage(committed); err != nil {
		abortErr := producer.AbortTxn()
		if abortErr != nil {
			return fmt.Errorf("send message failed: %w; abort failed: %v", err, abortErr)
		}
		return err
	}
	if err = producer.CommitTxn(); err != nil {
		return err
	}
	log.Println("Committed first transaction")

	// 事务 2：发送后主动回滚，消息对 read_committed 消费者不可见。
	if err = producer.BeginTxn(); err != nil {
		return err
	}

	aborted := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("k2"),
		Value: sarama.StringEncoder("aborted"),
	}
	if _, _, err = producer.SendMessage(aborted); err != nil {
		abortErr := producer.AbortTxn()
		if abortErr != nil {
			return fmt.Errorf("send aborted message failed: %w; abort failed: %v", err, abortErr)
		}
		return err
	}
	if err = producer.AbortTxn(); err != nil {
		return err
	}
	log.Println("Aborted second transaction")
	return nil
}

func consumeOneCommittedMessage(ctx context.Context, brokers []string, topic string, out chan<- string) (err error) {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V3_6_0_0
	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	// 核心：只读取已提交事务的数据。
	consumerConfig.Consumer.IsolationLevel = sarama.ReadCommitted

	consumer, err := sarama.NewConsumer(brokers, consumerConfig)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := consumer.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := partitionConsumer.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-partitionConsumer.Messages():
			if !ok {
				return errors.New("partition consumer closed")
			}
			out <- string(msg.Value)
			return nil
		case consumeErr, ok := <-partitionConsumer.Errors():
			if !ok {
				return errors.New("partition consumer error channel closed")
			}
			return consumeErr
		}
	}
}
