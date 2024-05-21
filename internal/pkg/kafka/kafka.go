package kafka

import (
	"go-web-scaffold/internal/pkg/logging"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var Producer sarama.SyncProducer

func InitKafka(brokers []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	Producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logging.Logger.Fatal("Failed to start Kafka producer", zap.Error(err))
	}

	logging.Logger.Info("Kafka producer started successfully")
}

func SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := Producer.SendMessage(msg)
	if err != nil {
		logging.Logger.Error("Failed to send message to Kafka", zap.Error(err))
		return
	}

	logging.Logger.Info("Message sent to Kafka",
		zap.String("topic", topic),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset))
}
