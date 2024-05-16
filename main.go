package main

import (
	"github.com/sirupsen/logrus"
	"svc-activity/cmd"
	"svc-activity/utils"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(utils.PrintMessageWithError("Can't start application command!", err))
	}
}
