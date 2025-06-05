package main

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// //生产者

func newKafkaProducer(brokers []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}

	return producer
}
func main() {
	router := gin.Default()
	producer := newKafkaProducer([]string{"192.168.3.55:9092"})

	router.POST("/send", func(c *gin.Context) {
		message := c.PostForm("message")
		if message == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "message is required"})
			return
		}

		msg := &sarama.ProducerMessage{
			Topic: "testtesting",
			Value: sarama.StringEncoder(message),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"partition": partition, "offset": offset})
	})

	router.Run(":8080")
}
