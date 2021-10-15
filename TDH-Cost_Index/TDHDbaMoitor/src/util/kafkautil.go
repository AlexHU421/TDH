package util

import (
	"github.com/Shopify/sarama"
)

func ProduceSendMsg (msg string,producer sarama.SyncProducer,TopicInformation string) {
	producer.SendMessage(&sarama.ProducerMessage{Topic:TopicInformation,Key:nil,Value: sarama.StringEncoder(msg)})


}
