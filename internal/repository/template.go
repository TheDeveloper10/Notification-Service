package repository

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util/code"
)

type ITemplateRepository interface {
	Insert(template *entity.TemplateEntity) (int, code.StatusCode)
	Get(id int) (*entity.TemplateEntity, code.StatusCode)
	GetBulk(filter *entity.TemplateFilter) (*[]entity.TemplateEntity, code.StatusCode)
	Update(template *entity.TemplateEntity) code.StatusCode
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