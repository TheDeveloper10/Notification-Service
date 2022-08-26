package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
)

type IClientRepository interface {
	GetClient(*entity.ClientCredentials) *entity.ClientEntity
	GenerateAccessToken(*entity.ClientEntity) *entity.AccessToken
	GetClientFromAccessToken(*entity.AccessToken) (*entity.ClientEntity, int)
	CreateClient(*entity.ClientEntity) *entity.ClientCredentials
}

func NewClientRepository(isMock bool) IClientRepository {
	if isMock {
		return &impl.MockClientRepository{}
	} else {
		repo := impl.BasicClientRepository{}
		repo.Init()
		return &repo
	}
}
