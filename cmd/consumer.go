package cmd

import (
	"context"
	"fmt"
	segmentioKafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
	"svc-activity/config"
	"sync"
	"syscall"
	"time"
)

var greetConsumer = `
     _        _   _       _ _            ____                                          
    / \   ___| |_(_)_   _(_) |_ _   _   / ___|___  _ __  ___ _   _ _ __ ___   ___ _ __ 
   / _ \ / __| __| \ \ / / | __| | | | | |   / _ \| '_ \/ __| | | | '_ ' _ \ / _ \ '__|
  / ___ \ (__| |_| |\ V /| | |_| |_| | | |__| (_) | | | \__ \ |_| | | | | | |  __/ |
 /_/   \_\___|\__|_| \_/ |_|\__|\__, |  \____\___/|_| |_|___/\__,_|_| |_| |_|\___|_|
                                |___/
`

func RegisterConsumer(kafkaConfig config.KafkaConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "activity consumer",
		Run: func(cmd *cobra.Command, args []string) {

			// greet
			fmt.Println(greetConsumer)

			// worker
			var wg sync.WaitGroup
			for i := 0; i < kafkaConfig.WorkerPool; i++ {
				wg.Add(1)
				go startReaderActivity(kafkaConfig, &wg)
			}
			wg.Wait()

		},
	}
}

func startReaderActivity(kafkaConfig config.KafkaConfig, wg *sync.WaitGroup) {

	// activity
	reader := segmentioKafka.NewReader(segmentioKafka.ReaderConfig{
		Brokers: strings.Split(kafkaConfig.BootstrapServers, ","),
		Topic:   kafkaConfig.Topic.Activity,
		GroupID: kafkaConfig.GroupID,
	})

	// start consumer
	go startConsumer(reader)

	// wait terminate signal
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM)
	<-termSignal

	// close consumer
	closeConsumer(reader)

	// finish
	wg.Done()
}

func startConsumer(reader *segmentioKafka.Reader) {
	logrus.Info("starting consumer..")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(fmt.Sprintf("consume error on topic %s", msg.Topic), err)
			time.Sleep(time.Second) // prevent massive error when error occurred
		} else {
			logrus.Info(fmt.Sprintf("message received at topic %s offset %d key %s partition %d", msg.Topic, msg.Offset, string(msg.Key), msg.Partition))
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
