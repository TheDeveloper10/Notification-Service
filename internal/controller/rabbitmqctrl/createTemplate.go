package rabbitmqctrl

import (
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
)
func NewCreateTemplateV1Controller(templateRepository repository.ITemplateRepository) iface.IRabbitMQController {
	return &basicCreateTemplateV1Controller{
		templateRepository,
	}
}

type basicCreateTemplateV1Controller struct {
	templateRepository repository.ITemplateRepository
}

func (bctc *basicCreateTemplateV1Controller) QueueName() string {
	return util.Config.RabbitMQ.TemplatesQueueName
}

func (bctc *basicCreateTemplateV1Controller) QueueCapacity() int {
	return util.Config.RabbitMQ.TemplatesQueueMax
}

func (bctc *basicCreateTemplateV1Controller) Handle(bytes []byte) (any, bool) {
	reqObj := dto.CreateTemplateRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return util.ErrorListFromTextError("Failed to parse JSON"), true
	}

	templateEntity := reqObj.ToEntity()
	id, status := bctc.templateRepository.Insert(templateEntity)
	if status != code.StatusSuccess {
		return nil, false
	}

	return id, true
}