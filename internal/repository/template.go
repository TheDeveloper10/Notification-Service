package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util"
)

type ITemplateRepository interface {
	Insert(entity *entity.TemplateEntity) (int, util.RepoStatusCode)
	Get(id int) (*entity.TemplateEntity, util.RepoStatusCode)
	GetBulk(filter *entity.TemplateFilter) (*[]entity.TemplateEntity, util.RepoStatusCode)
	Update(entity *entity.TemplateEntity) util.RepoStatusCode
	Delete(id int) util.RepoStatusCode
}

// ----------------------------------
// Template Repository Factories
// ----------------------------------

func NewTemplateRepository() ITemplateRepository {
	repo := impl.BasicTemplateRepository{}
	repo.Init()
	return &repo
}

func NewMockTemplateRepository() ITemplateRepository {
	return &impl.MockTemplateRepository{}
}