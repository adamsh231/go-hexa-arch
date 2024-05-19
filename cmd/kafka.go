package cmd

import (
	"fmt"
	"svc-activity/config"
	"svc-activity/internal/adapter/handler/consumer"
	"svc-activity/internal/adapter/libraries/pubsub"
	"svc-activity/utils"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var greetConsumer = `
 _  __    _    _____ _  __    _       ____ __  __ ____  
| |/ /   / \  |  ___| |/ /   / \     / ___|  \/  |  _ \ 
| ' /   / _ \ | |_  | ' /   / _ \   | |   | |\/| | | | |
| . \  / ___ \|  _| | . \  / ___ \  | |___| |  | | |_| |
|_|\_\/_/   \_\_|   |_|\_\/_/   \_\  \____|_|  |_|____/ 
														
`

func RegisterConsumer() *cobra.Command {
	return &cobra.Command{
		Use:   "kafka",
		Short: "kafka listener",
		Run: func(cmd *cobra.Command, args []string) {

			// greet
			fmt.Println(greetConsumer)

			// start consume
			consume()

		},
	}
}

type topicHandler struct {
	Topic   string
	Handler func(msg []byte)
}

func consume() {

	var wg sync.WaitGroup

	// setup config
	getConfig, err := config.SetupConfig()
	if err != nil {
		logrus.Fatal(utils.PrintMessageWithError("failed to set up configuration", err))
	}
	kafkaConfig := getConfig.Kafka

	// init injector
	inject := config.InitInjection(getConfig)
	handler := consumer.NewHandler(inject)

	// register topicHandlers - topic and handler added here and would be automatically has independent consumer
	topicHandlers := []topicHandler{
		{
			Topic:   kafkaConfig.Topic.Activity,
			Handler: handler.ReceiveAndInsertActivity,
		},
	}

	// running n worker consumer per topic
	for _, topicHandler := range topicHandlers {
		kafkaLib := pubsub.NewKafkaLibrary(kafkaConfig.BootstrapServers, topicHandler.Topic, kafkaConfig.GroupID)
		dispatchWorker(&wg, kafkaConfig.WorkerPool, kafkaLib.Subscribe, topicHandler.Handler)
	}

	wg.Wait()

	// close config
	getConfig.CloseConfig()
}

func dispatchWorker(wg *sync.WaitGroup, workerCount int, job func(wg *sync.WaitGroup, handler func(message []byte)), handler func(message []byte)) {
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go job(wg, handler)
	}
}
