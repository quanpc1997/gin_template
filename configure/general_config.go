package configure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configure struct {
	KafkaBootstrapServer string
	KafkaPort            string
	KafkaTopic           string
	KafkaGroup           string
}

func (c *Configure) InitConfigure() *Configure {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.KafkaBootstrapServer = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	c.KafkaPort = os.Getenv("KAFKA_PORT")
	c.KafkaTopic = os.Getenv("KAFKA_TOPIC")
	c.KafkaGroup = os.Getenv("KAFKA_GROUP")
	return c
}
