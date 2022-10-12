package impl

import "notification-service/internal/entity"

type MockClientRepository struct {}

func (mcr *MockClientRepository) GetClient(credentials *entity.ClientCredentials) *entity.ClientEntity {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}
}

func (mcr *MockClientRepository) UpdateClient(clientID *string, client *entity.ClientEntity) int {
	if *clientID == "aa" {
		return 0
	} else if *clientID == "bb" {
		return 1
	}
	return 2
}

func (mcr *MockClientRepository) DeleteClient(clientID *string) int {
	if *clientID == "aa" {
		return 0
	} else if *clientID == "bb" {
		return 1
	}
	return 2
}

func (mcr *MockClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) *entity.AccessToken {
	return &entity.AccessToken{
		AccessToken: "123",
	}
}

func (mcr *MockClientRepository) GetClientFromAccessToken(accessToken *entity.AccessToken) (*entity.ClientEntity, int) {
	return &entity.ClientEntity{
		Permissions: entity.PermissionAll,
	}, 0
}

func (mcr *MockClientRepository) CreateClient(clientEntity *entity.ClientEntity) *entity.ClientCredentials {
	return &entity.ClientCredentials{
		Id: "1234",
		Secret: "Real Secret",
	}
}