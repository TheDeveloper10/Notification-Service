package common

import (
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
)

func ToNotificationEntity(
						targetId int,
						template *entity.TemplateEntity,
						message *string,
						request *dto.SendNotificationRequest,
						errs *[]dto.SendNotificationErrorData) *entity.NotificationEntity {
	replaced, err := FillPlaceholders(*message, request.Targets[targetId].Placeholders)
	if err != nil {
		*errs = append(*errs, dto.SendNotificationErrorData{
			TargetId: targetId,
			Messages: []string { err.Error() },
		})
		return nil
	}

	unfilledPlaceholders := GetPlaceholders(replaced)

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
