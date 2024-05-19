package main

import (
	"go-hexa/cmd"
	"go-hexa/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(utils.PrintMessageWithError("Can't start application command!", err))
	}
}
