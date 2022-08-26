package impl

import "notification-service/internal/entity"

type MockTemplateRepository struct {}

func (mtr *MockTemplateRepository) Insert(entity *entity.TemplateEntity) int {
	return 0
}

func (mtr *MockTemplateRepository) Get(id int) (*entity.TemplateEntity, int) {
	return &entity.TemplateEntity{
		Id: 1,
		ContactType: "email",
		Template: "Hi, @{firstName}!",
		Language: "EN",
		Type: "test",
	}, 0
}

func (mtr *MockTemplateRepository) GetBulk(filter *entity.TemplateFilter) *[]entity.TemplateEntity {
	return &[]entity.TemplateEntity{
		{
			Id: 5,
			ContactType: "push",
			Template: "Hi, @{firstName}!",
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