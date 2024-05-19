package pubsub

import (
	"context"
	"fmt"
	"go-hexa/internal/core/port/libraries"
	"go-hexa/utils"
	"strings"
	"sync"
	"time"

	segmentioKafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type kafkaLibrary struct {
	bootstrapServers string
	topic            string
	groupID          string
}

func NewKafkaLibrary(servers, topic, groupID string) libraries.IPubSubLibrary {
	return kafkaLibrary{
		bootstrapServers: servers,
		topic:            topic,
		groupID:          groupID,
	}
}

func (lib kafkaLibrary) Publish() {
}

func (lib kafkaLibrary) Subscribe(wg *sync.WaitGroup, handler func(message []byte)) {

	// activity
	reader := segmentioKafka.NewReader(segmentioKafka.ReaderConfig{
		Brokers: strings.Split(lib.bootstrapServers, ","),
		Topic:   lib.topic,
		GroupID: lib.groupID,
	})

	// start consumer
	go startConsumer(reader, handler)

	// terminate signal
	utils.WaitTerminateSignal()

	// close consumer
	closeConsumer(reader)

	wg.Done()
}

func startConsumer(reader *segmentioKafka.Reader, handler func(message []byte)) {
	logrus.Info("starting consumer..")
	for {

		// receive message
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(fmt.Sprintf("consume error on topic %s", msg.Topic), err)
			time.Sleep(time.Second) // prevent massive error when error occurred
		} else {
			logrus.Info(fmt.Sprintf("message received at topic %s offset %d key %s partition %d", msg.Topic, msg.Offset, string(msg.Key), msg.Partition))
			handler(msg.Value)
		}

	}
}

func closeConsumer(reader *segmentioKafka.Reader) {
	logrus.Info(fmt.Sprintf("signal received, closing consumer..."))
	if err := reader.Close(); err != nil {
		logrus.Error("failed to close reader", err)
	}
	logrus.Info(fmt.Sprintf("consumer closed"))
}
