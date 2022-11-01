package repository

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util"
)

type ITemplateRepository interface {
	Insert(entity *entity2.TemplateEntity) (int, util.RepoStatusCode)
	Get(id int) (*entity2.TemplateEntity, util.RepoStatusCode)
	GetBulk(filter *entity2.TemplateFilter) (*[]entity2.TemplateEntity, util.RepoStatusCode)
	Update(entity *entity2.TemplateEntity) util.RepoStatusCode
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