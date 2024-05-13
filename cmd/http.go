package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var greetHTTP = `
     _        _   _       _ _           _   _ _____ _____ ____  
    / \   ___| |_(_)_   _(_) |_ _   _  | | | |_   _|_   _|  _ \ 
   / _ \ / __| __| \ \ / / | __| | | | | |_| | | |   | | | |_) |
  / ___ \ (__| |_| |\ V /| | |_| |_| | |  _  | | |   | | |  __/ 
 /_/   \_\___|\__|_| \_/ |_|\__|\__, | |_| |_| |_|   |_| |_|    
                                |___/                           
`

func RegisterHTTP() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "http application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(greetHTTP)
		},
	}
}
