package main

import (
	"log"
	"svc-activity/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("Can't start application command!")
	}
}
