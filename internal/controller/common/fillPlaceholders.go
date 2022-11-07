package common

import (
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
	"strings"
)

func FillPlaceholders(text string, placeholders []dto.TemplatePlaceholder) (*string, error) {
	for i := 0; i < len(placeholders); i++ {
		placeholder := &(placeholders[i])
		if err := placeholder.Validate(); err != nil {
			return nil, err
		}

		key := "@{" + placeholder.Key + "}"
		text = strings.ReplaceAll(text, key, placeholder.Value)
	}

	return &text, nil
}

func FillPlaceholdersOnTemplate(template *entity.TemplateEntity, placeholders []dto.TemplatePlaceholder) error {
	if template.Body.Email != nil {
		edited, err := FillPlaceholders(*template.Body.Email, placeholders)
		if util.ManageError(err) {
			return err
		}
		template.Body.Email = edited
	}

	if template.Body.SMS != nil {
		edited, err := FillPlaceholders(*template.Body.SMS, placeholders)
		if util.ManageError(err) {
			return err
		}
		template.Body.SMS = edited
	}

	if template.Body.Push != nil {
		edited, err := FillPlaceholders(*template.Body.Push, placeholders)
		if util.ManageError(err) {
			return err
		}
		template.Body.Push = edited
	}

	return nil
}