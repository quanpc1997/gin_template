package configure

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	config *Configure
}

func (kp *KafkaConsumer) InitKafkaConsumer() {
	kafka_bootstrap_server := kp.config.KafkaBootstrapServer + kp.config.KafkaPort
	// Tạo reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafka_bootstrap_server},
		Topic:     kp.config.KafkaTopic,
		GroupID:   kp.config.KafkaGroup,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error while reading message: %v", err)
		}
		fmt.Printf("Received message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))
	}
}
