package impl

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util/code"
)

type MockClientRepository struct {}

func (mcr *MockClientRepository) GetClient(credentials *entity.ClientCredentials) (*entity.ClientEntity, code.StatusCode) {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}, code.StatusSuccess
}

func (mcr *MockClientRepository) UpdateClient(clientID *string, client *entity.ClientEntity) code.StatusCode {
	if *clientID == "aa" {
		return code.StatusSuccess
	} else if *clientID == "bb" {
		return code.StatusNotFound
	}
	return code.StatusError
}

func (mcr *MockClientRepository) DeleteClient(clientID *string) code.StatusCode {
	if *clientID == "aa" {
		return code.StatusSuccess
	} else if *clientID == "bb" {
		return code.StatusNotFound
	}
	return code.StatusError
}

func (mcr *MockClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) (*entity.AccessToken, code.StatusCode) {
	return &entity.AccessToken{
		AccessToken: "123",
	}, code.StatusSuccess
}

func (mcr *MockClientRepository) GetClientFromAccessToken(accessToken *entity.AccessToken) (*entity.ClientEntity, code.StatusCode) {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}, code.StatusSuccess
}

func (mcr *MockClientRepository) CreateClient(clientEntity *entity.ClientEntity) (*entity.ClientCredentials, code.StatusCode) {
	return &entity.ClientCredentials{
		Id: "1234",
		Secret: "Real Secret",
	}, code.StatusSuccess
}