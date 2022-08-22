package repository

import (
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"

	log "github.com/sirupsen/logrus"
)

type ClientRepository interface {
	GetClient(*entity.ClientCredentials) *entity.ClientEntity
	GenerateAccessToken(*entity.ClientEntity) *entity.AccessToken
	GetClientFromAccessToken(*entity.AccessToken) *entity.ClientEntity
}

type basicClientRepository struct {
	sg *util.StringGenerator
}

func NewClientRepository() ClientRepository {
	sg := util.StringGenerator{}
	sg.Init()
	return &basicClientRepository{
		sg: &sg,
	}
}

func (bcr *basicClientRepository) GetClient(credentials *entity.ClientCredentials) *entity.ClientEntity {
	rows := client.SQLClient.Query(
		"select Permissions from Clients where Id=? and Secret=?",
		credentials.Id, credentials.Secret,
	)
	if rows == nil {
		return nil
	}

	if rows.Next() {
		record := entity.ClientEntity{}
		err3 := rows.Scan(&record.Permissions)
		if helper.IsError(err3) {
			return nil
		}

		log.Info("Fetched client with id " + credentials.Id)
		return &record
	}

	return nil
}

func (bcr *basicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) *entity.AccessToken {
	token := bcr.sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into AccessTokens(AccessToken, Permissions) values(?, ?)",
		token, clientEntity.Permissions,
	)
	if res == nil {
		return nil
	}

	log.Info("Generated a new access token")
	return &entity.AccessToken{
		AccessToken: token,
	}
}

func (bcr *basicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) *entity.ClientEntity {
	// Perhaps replace the memory table with a Redis Cache

	rows := client.SQLClient.Query("select Permissions from AccessTokens where AccessToken=?", token.AccessToken)

	if rows.Next() {
		record := entity.ClientEntity{}
		err3 := rows.Scan(&record.Permissions)
		if helper.IsError(err3) {
			return nil
		}

		log.Info("Fetched client with from access token " + token.AccessToken)
		return &record
	}

	return nil
}
