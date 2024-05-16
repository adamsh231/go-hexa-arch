package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitTerminateSignal(){
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM)
	<-termSignal
}