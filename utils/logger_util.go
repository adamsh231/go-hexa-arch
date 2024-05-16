package utils

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
)

func LogrusWithPayload(payload string) *logrus.Entry {
	return logrus.WithField("payload", payload)
}

func APMStartTransaction(name string) *apm.Transaction{
	return apm.DefaultTracer().StartTransaction(name, "request")
}