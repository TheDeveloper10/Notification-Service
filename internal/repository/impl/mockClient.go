package impl

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
)

type MockClientRepository struct {}

func (mcr *MockClientRepository) GetClient(credentials *entity.ClientCredentials) (*entity.ClientEntity, util.RepoStatusCode) {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}, util.RepoStatusSuccess
}

func (mcr *MockClientRepository) UpdateClient(clientID *string, client *entity.ClientEntity) util.RepoStatusCode {
	if *clientID == "aa" {
		return util.RepoStatusSuccess
	} else if *clientID == "bb" {
		return util.RepoStatusNotFound
	}
	return util.RepoStatusError
}

func (mcr *MockClientRepository) DeleteClient(clientID *string) util.RepoStatusCode {
	if *clientID == "aa" {
		return util.RepoStatusSuccess
	} else if *clientID == "bb" {
		return util.RepoStatusNotFound
	}
	return util.RepoStatusError
}

func (mcr *MockClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) (*entity.AccessToken, util.RepoStatusCode) {
	return &entity.AccessToken{
		AccessToken: "123",
	}, util.RepoStatusSuccess
}

func (mcr *MockClientRepository) GetClientFromAccessToken(accessToken *entity.AccessToken) (*entity.ClientEntity, util.RepoStatusCode) {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}, util.RepoStatusSuccess
}

func (mcr *MockClientRepository) CreateClient(clientEntity *entity.ClientEntity) (*entity.ClientCredentials, util.RepoStatusCode) {
	return &entity.ClientCredentials{
		Id: "1234",
		Secret: "Real Secret",
	}, util.RepoStatusSuccess
}