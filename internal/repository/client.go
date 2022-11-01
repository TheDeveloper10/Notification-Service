package repository

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util"
)

type IClientRepository interface {
	CreateClient(*entity.ClientEntity) (*entity.ClientCredentials, util.RepoStatusCode)
	GetClient(*entity.ClientCredentials) (*entity.ClientEntity, util.RepoStatusCode)
	UpdateClient(*string, *entity.ClientEntity) util.RepoStatusCode
	DeleteClient(*string) util.RepoStatusCode

	GenerateAccessToken(*entity.ClientEntity) (*entity.AccessToken, util.RepoStatusCode)
	GetClientFromAccessToken(*entity.AccessToken) (*entity.ClientEntity, util.RepoStatusCode)
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