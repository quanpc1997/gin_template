package configure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
}

func (kp KafkaProducer) InitKafkaProducer() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{}, // Cân bằng load theo kích thước message
	})
	defer writer.Close()

	for i := 0; i < 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprintf("Hello Kafka %d", i)),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf("failed to write message: %v", err)
		}
		fmt.Printf("Produced message: %s\n", msg.Value)
		time.Sleep(time.Second)
	}
}
