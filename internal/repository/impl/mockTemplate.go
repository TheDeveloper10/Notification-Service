package impl

import "notification-service/internal/entity"

type MockTemplateRepository struct {}

func (mtr *MockTemplateRepository) Insert(entity *entity.TemplateEntity) int {
	return 0
}

func (mtr *MockTemplateRepository) Get(id int) (*entity.TemplateEntity, int) {
	template := "Hi, @{firstName}!"
	return &entity.TemplateEntity{
		Id: 1,
		Body: entity.TemplateBody{
			Email: &template,
		},
		Language: "EN",
		Type: "test",
	}, 0
}

func (mtr *MockTemplateRepository) GetBulk(filter *entity.TemplateFilter) *[]entity.TemplateEntity {
	template := "Hi, @{firstName}!"
	return &[]entity.TemplateEntity{
		{
			Id: 5,
			Body: entity.TemplateBody{
				Push: &template,
			},
			Language: "EN",
			Type: "test",
		},
	}
}

func (mtr *MockTemplateRepository) Update(templateEntity *entity.TemplateEntity) int {
	return 0
}

func (mtr *MockTemplateRepository) Delete(id int) bool {
	return true
}