package impl

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/util/code"
)

type MockTemplateRepository struct {}

func (mtr *MockTemplateRepository) Insert(entity *entity2.TemplateEntity) (int, code.StatusCode) {
	return 0, code.StatusSuccess
}

func (mtr *MockTemplateRepository) Get(id int) (*entity2.TemplateEntity, code.StatusCode) {
	if id == 4 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				Email: &template,
			},
			Language: "EN",
			Type:     "test",
		}, code.StatusSuccess
	} else if id == 3 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				SMS: &template,
			},
			Language: "EN",
			Type:     "test",
		}, code.StatusSuccess
	} else if id == 2 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				Push: &template,
			},
			Language: "EN",
			Type:     "test",
		}, code.StatusSuccess
	} else if id == 1 {
		return nil, code.StatusNotFound
	}

	return nil, code.StatusError
}

func (mtr *MockTemplateRepository) GetBulk(filter *entity2.TemplateFilter) (*[]entity2.TemplateEntity, code.StatusCode) {
	template := "Hi, @{firstName}!"
	return &[]entity2.TemplateEntity{
		{
			Id: 5,
			Body: entity2.TemplateBody{
				Push: &template,
			},
			Language: "EN",
			Type: "test",
		},
	}, code.StatusSuccess
}

func (mtr *MockTemplateRepository) Update(templateEntity *entity2.TemplateEntity) code.StatusCode {
	return code.StatusSuccess
}

func (mtr *MockTemplateRepository) Delete(id int) code.StatusCode {
	return code.StatusSuccess
}