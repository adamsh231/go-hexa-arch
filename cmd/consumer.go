package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"svc-activity/config"
	"svc-activity/internal/adapter/handler/consumer"
	"svc-activity/libraries"
	"sync"
)

var greetConsumer = `
     _        _   _       _ _            ____                                          
    / \   ___| |_(_)_   _(_) |_ _   _   / ___|___  _ __  ___ _   _ _ __ ___   ___ _ __ 
   / _ \ / __| __| \ \ / / | __| | | | | |   / _ \| '_ \/ __| | | | '_ ' _ \ / _ \ '__|
  / ___ \ (__| |_| |\ V /| | |_| |_| | | |__| (_) | | | \__ \ |_| | | | | | |  __/ |
 /_/   \_\___|\__|_| \_/ |_|\__|\__, |  \____\___/|_| |_|___/\__,_|_| |_| |_|\___|_|
                                |___/
`

func RegisterConsumer() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "activity consumer",
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
		log.Fatal(fmt.Sprintf("Failed to set up configuration err: %v", err.Error()))
	}
	kafkaConfig := getConfig.Kafka

	// init injector
	inject := config.InitInjection(getConfig)
	handler := consumer.NewHandler(inject)


	// register topicHandlers - topic and handler added here and would be automatically consume
	topicHandlers := []topicHandler{
		{
			Topic:   kafkaConfig.Topic.Activity,
			Handler: handler.ReceiveAndInsertActivity,
		},
	}

	// running n worker consumer per topic
	for _, topicHandler := range topicHandlers {
		kafkaLib := libraries.NewKafkaLibrary(kafkaConfig.BootstrapServers, topicHandler.Topic, kafkaConfig.GroupID)
		dispatchWorker(&wg, kafkaConfig.WorkerPool, kafkaLib.Consume, topicHandler.Handler)
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
