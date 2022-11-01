package impl

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/util"
)

type MockTemplateRepository struct {}

func (mtr *MockTemplateRepository) Insert(entity *entity2.TemplateEntity) (int, util.RepoStatusCode) {
	return 0, util.RepoStatusSuccess
}

func (mtr *MockTemplateRepository) Get(id int) (*entity2.TemplateEntity, util.RepoStatusCode) {
	if id == 4 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				Email: &template,
			},
			Language: "EN",
			Type:     "test",
		}, util.RepoStatusSuccess
	} else if id == 3 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				SMS: &template,
			},
			Language: "EN",
			Type:     "test",
		}, util.RepoStatusSuccess
	} else if id == 2 {
		template := "Hi, @{firstName}!"
		return &entity2.TemplateEntity{
			Id: 1,
			Body: entity2.TemplateBody{
				Push: &template,
			},
			Language: "EN",
			Type:     "test",
		}, util.RepoStatusSuccess
	} else if id == 1 {
		return nil, util.RepoStatusNotFound
	}

	return nil, util.RepoStatusError
}

func (mtr *MockTemplateRepository) GetBulk(filter *entity2.TemplateFilter) (*[]entity2.TemplateEntity, util.RepoStatusCode) {
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
	}, util.RepoStatusSuccess
}

func (mtr *MockTemplateRepository) Update(templateEntity *entity2.TemplateEntity) util.RepoStatusCode {
	return util.RepoStatusSuccess
}

func (mtr *MockTemplateRepository) Delete(id int) util.RepoStatusCode {
	return util.RepoStatusSuccess
}