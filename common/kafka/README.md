# Kafka 事务性生产者示例（Sarama）

核心演示代码在 `transactional_producer.go`：使用 Sarama 在一个事务中向四个 topic 发送消息（更新订单状态、扣减库存、发送通知、写入审计日志）。

使用要点：
- Kafka 版本需 >= 0.11；推荐 >= 2.5（示例配置为 `V2_5_0_0`）。
- 生产者必须设置：`Idempotent=true`、`RequiredAcks=WaitForAll`、`Return.Successes=true`、`Net.MaxOpenRequests=1`。
- 设置稳定且全局唯一的 `transactional.id`（示例中通过 `TransactionConfig{ID: ...}`）。
- 开启事务：`BeginTxn()`，失败则 `AbortTxn()`，成功则 `CommitTxn()`。
- 消费端读取提交后的记录需配置 `isolation.level=read_committed`。

简单调用示例：

```go
ctx := context.Background()
brokers := []string{"localhost:9092"}
topics := [4]string{"update_order_status", "deduct_inventory", "send_notification", "write_audit_log"}
payloads := [4]string{"status:paid", "sku:123 qty:1", "notify:user:ok", "audit:ok"}
err := kafka.ProducePaymentSuccessTransaction(ctx, brokers, "txn-producer-1", "order-123", "order-123", payloads, topics)
if err != nil { log.Fatal(err) }
```

注意：`transactional.id` 应与实例生命周期绑定且保持稳定；不要在同一时刻被多个生产者实例复用。


