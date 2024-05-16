package utils

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type LogUtil struct {
	id string
}

func NewLogUtil() LogUtil {
	return LogUtil{id: uuid.NewString()}
}

// LogrusWithPayload use this if you want to track every log with uuid
func (util LogUtil) LogrusWithPayload(payload string) *logrus.Entry {
	return logrus.WithField("payload", payload).WithField("id", util.id)
}

// Logrus use this if you want to track every log with uuid
func (util LogUtil) Logrus(payload string) *logrus.Entry {
	return logrus.WithField("id", util.id)
}
