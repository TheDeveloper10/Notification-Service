package httpctrl

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/common"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
)

func NewNotificationV1Controller(templateRepository repository.ITemplateRepository,
								NotificationRepository repository.INotificationRepository,
								clientRepository repository.IClientRepository) iface.IHTTPController {
	return &basicNotificationV1Controller{
		common.NotificationSender{
			TemplateRepository:     templateRepository,
			NotificationRepository: NotificationRepository,
		},
		clientRepository,
	}
}

type basicNotificationV1Controller struct {
	common.NotificationSender
	clientRepository       repository.IClientRepository
}

func (bnc *basicNotificationV1Controller) CreateRoutes(router *rem.Router) {
	router.
		NewRoute("/v1/notifications").
		Get(bnc.getBulk).
		Post(bnc.send)
}

func (bnc *basicNotificationV1Controller) getBulk(res rem.IResponse, req rem.IRequest) bool {
	// GET /notifications
	// GET /notifications?page=24 (size = default = 20)
	// GET /notifications?size=50 (page = default = 1)
	// GET /notifications?appId=aa-bb
	// GET /notifications?templateId=45
	// GET /notifications?startTime=17824254
	// GET /notifications?endTime=17824254
	if !layer.AccessTokenMiddleware(bnc.clientRepository, res, req, entity.PermissionReadSentNotifications) {
		return true
	}

	filter := entity.NotificationFilterFromRequest(req, res)
	if filter == nil {
		return true
	}

	notifications, status := bnc.NotificationRepository.GetBulk(filter)
	if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get anything"))
	} else {
		res.Status(http.StatusOK).JSON(*notifications)
	}

	return true
}

func (bnc *basicNotificationV1Controller) send(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(bnc.clientRepository, res, req, entity.PermissionSendNotifications) {
		return true
	}

	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	value, _, status := bnc.NotificationSender.Send(&reqObj)
	if value == nil {
		res.Status(status)
	} else if v, ok := value.(error); ok {
		res.Status(status).JSON(util.ErrorListFromTextError(v.Error()))
	} else {
		res.Status(status).JSON(value)
	}

	return true
}
