package consumer

import (
	"encoding/json"
	"go-hexa/config"
	"go-hexa/internal/core/domain/entities"
	"go-hexa/utils"
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
		utils.NewLogUtil().LogrusWithPayload(string(message)).Error(err)
	}

	// handling
	_ = handler.injector.ActivityService.InsertActivity(activityInput)
}
