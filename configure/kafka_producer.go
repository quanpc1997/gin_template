package configure

import (
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaProducer struct {
	config *Configure
}

func (kp *KafkaProducer) InitKafkaProducer() {
	kafka_bootstrap_server := kp.config.KafkaBootstrapServer + kp.config.KafkaPort

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafka_bootstrap_server},
		Topic:    kp.config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	for i := 0; i < 10; i++ {
		zap.L().Info("HAHA")
		// msg := kafka.Message{
		// 	Key:   []byte(uuid.New().String()),
		// 	Value: []byte(fmt.Sprintf("Hello Kafka %d", i)),
		// }

		// err := writer.WriteMessages(context.Background(), msg)
		// if err != nil {
		// 	log.Fatalf("failed to write message: %v", err)
		// }
		// fmt.Printf("Produced message: %s\n", msg.Value)
		time.Sleep(time.Second)
	}
}
