package repository

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util/code"
)

type ITemplateRepository interface {
	Insert(entity *entity2.TemplateEntity) (int, code.StatusCode)
	Get(id int) (*entity2.TemplateEntity, code.StatusCode)
	GetBulk(filter *entity2.TemplateFilter) (*[]entity2.TemplateEntity, code.StatusCode)
	Update(entity *entity2.TemplateEntity) code.StatusCode
	Delete(id int) code.StatusCode
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