package rabbitmqctrl

import (
	"notification-service/internal/controller/common"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

func NewCreateNotificationV1Controller(TemplateRepository repository.ITemplateRepository,
									NotificationRepository repository.INotificationRepository) iface.IRabbitMQController {
	return &basicCreateNotificationV1Controller{
		common.NotificationSender{
			TemplateRepository:     TemplateRepository,
			NotificationRepository: NotificationRepository,
		},
	}
}

type basicCreateNotificationV1Controller struct {
	common.NotificationSender
}

func (bcnc *basicCreateNotificationV1Controller) QueueName() string {
	return util.Config.RabbitMQ.NotificationsQueueName
}

func (bcnc *basicCreateNotificationV1Controller) QueueCapacity() int {
	return util.Config.RabbitMQ.NotificationsQueueMax
}

func (bcnc *basicCreateNotificationV1Controller) Handle(bytes []byte) (any, bool) {
	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return util.ErrorListFromTextError("Failed to parse JSON"), true
	} else {
		validationErrs := reqObj.Validate()
		if validationErrs.ErrorsCount() > 0 {
			return validationErrs, true
		}
	}

	value, acknowledge, _ := bcnc.Send(&reqObj)
	if value == nil {
		return nil, acknowledge
	} else if v, ok := value.(error); ok {
		return util.ErrorListFromTextError(v.Error()), acknowledge
	} else {
		return value, acknowledge
	}
}
