package impl

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util/code"
)

type MockClientRepository struct {}

func (mcr *MockClientRepository) GetClient(credentials *entity.ClientCredentials) (*entity.ClientEntity, code.StatusCode) {
	return &entity.ClientEntity{
		ClientId: credentials.Id,
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

func (mcr *MockClientRepository) CreateClient(clientEntity *entity.ClientEntity) (*entity.ClientCredentials, code.StatusCode) {
	return &entity.ClientCredentials{
		Id: "1234",
		Secret: "Real Secret",
	}, code.StatusSuccess
}

func (mcr *MockClientRepository) VerifyToken(token *string, secret *string) code.StatusCode {
	return code.StatusSuccess
}

func (mcr *MockClientRepository) GenerateToken(clientEntity *entity.ClientEntity, secret *string, expiry int) (*string, code.StatusCode) {
	return nil, code.StatusSuccess
}

func (mcr *MockClientRepository) ExtractClientFromToken(token *string, secret *string) (*entity.ClientEntity, code.StatusCode) {
	return nil, code.StatusSuccess
}
