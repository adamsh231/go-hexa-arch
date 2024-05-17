package entities

import "time"

type InsertActivityInput struct {
	Service   string
	Version   string
	Message   string
	Activity  string
	CreatedBy string
	Data      map[string]interface{}
	Response  map[string]interface{}
}

type SearchActivityInput struct {
	Service string
	Created time.Time
	Page    int
	Limit   int
}

type SearchActivityOutput struct {
	ID        string `json:"id"`
	Service   string `json:"service,omitempty"`
	Version   string `json:"version,omitempty"`
	Message   string `json:"message,omitempty"`
	Activity  string `json:"activity,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	Created   string `json:"created,omitempty"`
}

type FindActivityOutput struct {
	ID        string                 `json:"id"`
	Service   string                 `json:"service,omitempty"`
	Version   string                 `json:"version,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Activity  string                 `json:"activity,omitempty"`
	CreatedBy string                 `json:"created_by,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Response  map[string]interface{} `json:"response,omitempty"`
	Created   string                 `json:"created,omitempty"`
}
