package utils

import "github.com/sirupsen/logrus"

func LogrusWithPayload(payload string) *logrus.Entry {
	return logrus.WithField("payload", payload)
}
