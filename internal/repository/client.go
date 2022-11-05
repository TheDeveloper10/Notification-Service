package repository

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util/code"
)

type IClientRepository interface {
	CreateClient(*entity.ClientEntity) (*entity.ClientCredentials, code.StatusCode)
	GetClient(*entity.ClientCredentials) (*entity.ClientEntity, code.StatusCode)
	UpdateClient(*string, *entity.ClientEntity) code.StatusCode
	DeleteClient(*string) code.StatusCode

	GenerateAccessToken(*entity.ClientEntity) (*entity.AccessToken, code.StatusCode)
	GetClientFromAccessToken(*entity.AccessToken) (*entity.ClientEntity, code.StatusCode)
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