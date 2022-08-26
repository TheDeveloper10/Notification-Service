package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
)

type ITemplateRepository interface {
	Insert(entity *entity.TemplateEntity) int
	Get(id int) (*entity.TemplateEntity, int)
	GetBulk(filter *entity.TemplateFilter) *[]entity.TemplateEntity
	Update(entity *entity.TemplateEntity) int
	Delete(id int) bool
}

func NewTemplateRepository(isMock bool) ITemplateRepository {
	if isMock {
		return &impl.MockTemplateRepository{}
	} else {
		repo := impl.BasicTemplateRepository{}
		repo.Init()
		return &repo
	}
}
