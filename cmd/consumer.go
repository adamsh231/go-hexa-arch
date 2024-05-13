package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
		Short: "consumer application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(greetConsumer)
		},
	}
}
