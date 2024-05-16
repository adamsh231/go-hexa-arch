package utils

import "go.elastic.co/apm/v2"

func APMStartTransaction(name string) *apm.Transaction{
	return apm.DefaultTracer().StartTransaction(name, "request")
}