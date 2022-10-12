package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
)

type IClientRepository interface {
	GetClient(*entity.ClientCredentials) *entity.ClientEntity
	UpdateClient(*string, *entity.ClientEntity) int
	DeleteClient(*string) int

	GenerateAccessToken(*entity.ClientEntity) *entity.AccessToken
	GetClientFromAccessToken(*entity.AccessToken) (*entity.ClientEntity, int)
	CreateClient(*entity.ClientEntity) *entity.ClientCredentials
}

// ----------------------------------
// Client Repository Factories
// ----------------------------------


func NewClientRepository() IClientRepository {
	repo := impl.BasicClientRepository{}
	repo.Init()
	return &repo
}

func NewMockClientRepository() IClientRepository {
	return &impl.MockClientRepository{}
}