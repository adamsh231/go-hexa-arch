package models

import "time"

const ActivityCollection = "log-activity"

type InsertActivityModel struct {
	Service   string
	Version   string
	Message   string
	Activity  string
	CreatedBy string
	Data      map[string]interface{}
	Response  map[string]interface{}
	Created   time.Time
}
