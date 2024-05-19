package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Art ref  => https://edukits.co/text-art/
var greetDefault = `
 __  __       _          ____ __  __ ____  
|  \/  | __ _(_)_ __    / ___|  \/  |  _ \ 
| |\/| |/ _  | | '_ \  | |   | |\/| | | | |
| |  | | (_| | | | | | | |___| |  | | |_| |
|_|  |_|\__,_|_|_| |_|  \____|_|  |_|____/ 
										   
`

func Execute() (err error) {

	// default command
	var rootCmd = &cobra.Command{
		Use:   "main",
		Short: "This default command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(greetDefault)
		},
	}

	// register command
	rootCmd.AddCommand(RegisterHTTP())
	rootCmd.AddCommand(RegisterConsumer())

	return rootCmd.Execute()
}
