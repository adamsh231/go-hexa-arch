package models

import "time"

const ActivityCollection = "log-activity"

type ActivityModel struct {
	ID        string `bson:"_id"`
	Service   string
	Version   string
	Message   string
	Activity  string
	CreatedBy string
	Data      map[string]interface{}
	Response  map[string]interface{}
	Created   time.Time
}
