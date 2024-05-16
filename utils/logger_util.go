package utils

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
	"os"
	"os/signal"
	"syscall"
)

func LogrusWithPayload(payload string) *logrus.Entry {
	return logrus.WithField("payload", payload)
}

func APMStartTransaction(name string) *apm.Transaction{
	return apm.DefaultTracer().StartTransaction(name, "request")
}

func WaitTerminateSignal(){
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM)
	<-termSignal
}