package consumer

import (
	"fmt"
	"svc-activity/config"
	"svc-activity/internal/core/domain/entities"
)

type Handler struct {
	injector config.ServiceInjector
}

func NewHandler(injector config.ServiceInjector) Handler {
	return Handler{injector: injector}
}

func (handler Handler) ReceiveAndInsertActivity(message []byte) {
	handler.injector.ActivityService.InsertActivity(entities.InsertActivityInput{})
	fmt.Println("ini handler")
}
