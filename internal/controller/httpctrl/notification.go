package httpctrl

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
	"strconv"
	"sync"
)

func NewNotificationV1Controller(templateRepository repository.ITemplateRepository,
								notificationRepository repository.INotificationRepository,
								clientRepository repository.IClientRepository) iface.IHTTPController {
	return &basicNotificationV1Controller{
		templateRepository,
		notificationRepository,
		clientRepository,
	}
}

type basicNotificationV1Controller struct {
	templateRepository     repository.ITemplateRepository
	notificationRepository repository.INotificationRepository
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

	notifications, status := bnc.notificationRepository.GetBulk(filter)
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

	bnc.internalSend(&reqObj, res)
	return true
}

func (bnc *basicNotificationV1Controller) internalSend(reqObj *dto.SendNotificationRequest, res rem.IResponse) {
	templateEntity, status := bnc.templateRepository.Get(reqObj.TemplateID)
	if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something was wrong with the database. Try again"))
		return
	} else if status == code.StatusNotFound {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Template was not found!"))
		return
	}

	filledResult := bnc.fillPlaceholders(templateEntity, &reqObj.UniversalPlaceholders, res)
	if !filledResult {
		return
	}

	var wg sync.WaitGroup
	errs := make([]dto.SendNotificationErrorData, 0)

	successCount := 0
	failedCount := 0
	targetCount := len(reqObj.Targets)
	for currentTarget := 0; currentTarget < targetCount; currentTarget++ {
		target := &(reqObj.Targets[currentTarget])

		targetErrors := target.Validate()
		targetErrors.Merge(target.ValidateAgainstTemplate(templateEntity))
		if targetErrors != nil && targetErrors.ErrorsCount() > 0 {
			errs = append(errs, dto.SendNotificationErrorData{
				TargetId: currentTarget,
				Messages: *targetErrors.GetErrors(),
			})
			continue
		}


		if target.Email != nil {
			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Email, reqObj, &errs)
			if notification != nil {
				go bnc.sendNotification(*notification, target.Email, bnc.notificationRepository.SendEmail, &successCount, &failedCount, &wg)
			}
		}


		if target.PhoneNumber != nil {
			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.SMS, reqObj, &errs)
			if notification != nil {
				go bnc.sendNotification(*notification, target.PhoneNumber, bnc.notificationRepository.SendSMS, &successCount, &failedCount, &wg)
			}
		}


		if target.FCMRegistrationToken != nil {
			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Push, reqObj, &errs)
			if notification != nil {
				go bnc.sendNotification(*notification, target.FCMRegistrationToken, bnc.notificationRepository.SendPush, &successCount, &failedCount, &wg)
			}
		}
	}

	wg.Wait()

	if failedCount > 0 || len(errs) > 0 {
		res.Status(http.StatusBadRequest).JSON(dto.SendNotificationsError{
			Errors:                        errs,
			SuccessfullySentNotifications: successCount,
			FailedNotifications:           failedCount,
		})
	} else {
		res.Status(http.StatusCreated).Text(strconv.Itoa(targetCount) + " notification(s) have been sent successfully!")
	}
}


func (bnc *basicNotificationV1Controller) sendNotification(
											notification entity.NotificationEntity,
											contactInfo *string,
											outsourceNotification func(notificationEntity *entity.NotificationEntity)bool,
											successCount *int,
											failedCount *int,
											wg *sync.WaitGroup) {
	notification.ContactInfo = *contactInfo

	wg.Add(1)
	defer wg.Done()
	if outsourceNotification(&notification) {
		status := bnc.notificationRepository.Insert(&notification)
		if status == code.StatusSuccess {
			(*successCount)++
			return
		}
	}
	(*failedCount)++
}

func (bnc *basicNotificationV1Controller) toNotificationEntity(
											targetId int,
											template *entity.TemplateEntity,
											message *string,
											request *dto.SendNotificationRequest,
											errs *[]dto.SendNotificationErrorData) *entity.NotificationEntity {
	replaced, err := dto.FillPlaceholders(*message, &request.Targets[targetId].Placeholders)
	if err != nil {
		*errs = append(*errs, dto.SendNotificationErrorData{
			TargetId: targetId,
			Messages: []string { err.Error() },
		})
		return nil
	}

	unfilledPlaceholders := dto.GetPlaceholders(replaced)

	if len(unfilledPlaceholders) > 0 {
		*errs = append(*errs, dto.SendNotificationErrorData{
			TargetId: targetId,
			Messages: []string { "Unfilled placeholders: " + unfilledPlaceholders },
		})

		return nil
	}

	return &entity.NotificationEntity{
		TemplateID:  template.Id,
		AppID:       request.AppID,
		Title:       request.Title,
		Message:     *replaced,
	}
}

func (bnc *basicNotificationV1Controller) fillPlaceholders(
											template *entity.TemplateEntity,
											placeholders *[]dto.TemplatePlaceholder,
											res rem.IResponse) bool {
	if(template.Body.Email != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.Email, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.Email = edited
	}

	if(template.Body.SMS != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.SMS, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.SMS = edited
	}

	if(template.Body.Push != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.Push, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.Push = edited
	}

	return true
}
