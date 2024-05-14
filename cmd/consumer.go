package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"svc-activity/config"
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

			// handler
			handler()

		},
	}
}

func handler(){

	var wg sync.WaitGroup

	// setup config
	getConfig, err := config.SetupConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to set up configuration err: %v", err.Error()))
	}
	kafkaConfig := getConfig.Kafka

	// topics
	topics := []string{
		kafkaConfig.Topic.Activity,
	}

	// running n worker consumer per topic
	for _, topic := range topics{
		kafkaLib := libraries.NewKafkaLibrary(kafkaConfig.BootstrapServers, topic, kafkaConfig.GroupID)
		dispatchWorker(&wg, kafkaConfig.WorkerPool, kafkaLib.Consume)
	}

	wg.Wait()

	// close config
	getConfig.CloseConfig()

}

func dispatchWorker(wg *sync.WaitGroup, workerCount int, job func(wg *sync.WaitGroup)){
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go job(wg)
	}
}
