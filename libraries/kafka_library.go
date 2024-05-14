package libraries

import (
	"context"
	"fmt"
	segmentioKafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type KafkaLibrary struct {
	BootstrapServers string
	Topic            string
	GroupID          string
}

func NewKafkaLibrary(servers, topic, groupID string) KafkaLibrary {
	return KafkaLibrary{
		BootstrapServers: servers,
		Topic:            topic,
		GroupID:          groupID,
	}
}

func (lib KafkaLibrary) Consume(wg *sync.WaitGroup, handler func(message []byte)) {

	// activity
	reader := segmentioKafka.NewReader(segmentioKafka.ReaderConfig{
		Brokers: strings.Split(lib.BootstrapServers, ","),
		Topic:   lib.Topic,
		GroupID: lib.GroupID,
	})

	// start consumer
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM)
	go startConsumer(reader, handler)

	// wait terminate signal
	<-termSignal

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
		}else{
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
