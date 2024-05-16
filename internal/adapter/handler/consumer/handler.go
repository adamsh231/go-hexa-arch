package consumer

import (
	"encoding/json"
	"svc-activity/config"
	"svc-activity/internal/core/domain/entities"
	"svc-activity/utils"
)

type Handler struct {
	injector config.ServiceInjector
}

func NewHandler(injector config.ServiceInjector) Handler {
	return Handler{injector: injector}
}

func (handler Handler) ReceiveAndInsertActivity(message []byte) {

	// apm
	apmTx := utils.APMStartTransaction("ReceiveAndInsertActivity")
	defer apmTx.End()

	// destruct
	var activityInput entities.InsertActivityInput
	if err := json.Unmarshal(message, &activityInput); err != nil {
		utils.LogrusWithPayload(string(message)).Error(err)
	}

	// handling
	handler.injector.ActivityService.InsertActivity(activityInput)
}
