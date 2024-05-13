package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"svc-activity/config"
)

// Art ref  => https://edukits.co/text-art/
var greetDefault = `
  __  __    _       _  ___   ___       _    ____ _____ _____     _____ _______   __
 |  \/  |  / \     | |/ _ \ / _ \     / \  / ___|_   _|_ _\ \   / /_ _|_   _\ \ / /
 | |\/| | / _ \ _  | | | | | | | |   / _ \| |     | |  | | \ \ / / | |  | |  \ V / 
 | |  | |/ ___ \ |_| | |_| | |_| |  / ___ \ |___  | |  | |  \ V /  | |  | |   | |  
 |_|  |_/_/   \_\___/ \___/ \___/  /_/   \_\____| |_| |___|  \_/  |___| |_|   |_|  
`

func Execute() (err error) {

	// get config
	_, err = config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to set up configuration err: %v", err.Error()))
	}

	// default command
	var rootCmd = &cobra.Command{
		Use:   "majoo",
		Short: "Majoo default command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(greetDefault)
		},
	}

	// register http
	rootCmd.AddCommand(RegisterHTTP())

	// register consumer
	rootCmd.AddCommand(RegisterConsumer())

	return rootCmd.Execute()
}
