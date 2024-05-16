package entities

type InsertActivityInput struct {
	Service   string                 `json:"service,omitempty"`
	Version   string                 `json:"version,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Activity  string                 `json:"activity,omitempty"`
	CreatedBy string                 `json:"created_by,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Response  map[string]interface{} `json:"response,omitempty"`
}
