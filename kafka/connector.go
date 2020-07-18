package kafka

import (
	"context"
	"fmt"
	"strings"
	"time"


	"github.com/alecthomas/log4go"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConnector struct {
	brokers string
	topic string
	partition int
	log      log4go.Logger
	writer *kafka.Writer
	reader *kafka.Reader
}

func (kfk *KafkaConnector) Init(brokers string, topic string, partition int, log log4go.Logger) {
	kfk.brokers = brokers
	kfk.topic = topic
	kfk.partition = partition
	kfk.log = log
	kfk.writer = nil
	kfk.reader = nil
}

func (kfk *KafkaConnector) InitWriter() error {

	kfk.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(kfk.brokers,","),
		Topic:   kfk.topic,
		//Partition: kfk.partition,
		Balancer: &kafka.LeastBytes{},
	})

	return nil
}

func (kfk *KafkaConnector) InitReader() error {
	kfk.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(kfk.brokers,","),
		Topic:   kfk.topic,
		//Partition: kfk.partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		CommitInterval: time.Second,
	})
	return nil
}

func (kfk *KafkaConnector) GetMessage(msg []byte) kafka.Message {
	msg2 := &kafka.Message{Value:msg}
	return *msg2
}

func (kfk *KafkaConnector) ReadMessage() (kafka.Message, error) {
	m, err := kfk.reader.ReadMessage(context.Background())
	return m, err
}

func (kfk *KafkaConnector) WriteMessage(msg kafka.Message) error {
	if kfk.writer != nil {
		fmt.Printf("Using writer\n")
		return kfk.writer.WriteMessages(context.Background(),msg)
	}
	return fmt.Errorf("Invalid connection to kafka")
}

func (kfk *KafkaConnector) Close() {
	kfk.writer.Close()
	kfk.reader.Close()
}

func (kfk *KafkaConnector) CloseWriter() {
	kfk.writer.Close()
}

func (kfk *KafkaConnector) CloseReader() {
	kfk.reader.Close()
}
