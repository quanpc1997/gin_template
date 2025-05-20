package configure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configure struct {
	S3Bucket  string
	SecretKey string
}

func (c *Configure) InitConfigure() *Configure {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.S3Bucket = os.Getenv("S3_BUCKET")
	c.SecretKey = os.Getenv("SECRET_KEY")
	return c
}
