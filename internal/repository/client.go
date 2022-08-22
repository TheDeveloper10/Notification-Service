package repository

import (
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"time"

	log "github.com/sirupsen/logrus"
)

type ClientRepository interface {
	GetClient(*entity.ClientCredentials) *entity.ClientEntity
	GenerateAccessToken(*entity.ClientEntity) *entity.AccessToken
	GetClientFromAccessToken(*entity.AccessToken) (*entity.ClientEntity, int)
	CreateClient(*entity.ClientEntity) *entity.ClientCredentials
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
	defer helper.HandledClose(rows)

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
		"insert into AccessTokens(AccessToken, Permissions, ExpiryTime) values(?, ?, unix_timestamp() + ?)",
		token, clientEntity.Permissions, helper.Config.HTTPServer.AccessTokenExpiryTime,
	)
	if res == nil {
		return nil
	}

	log.Info("Generated a new access token")
	return &entity.AccessToken{
		AccessToken: token,
	}
}

func (bcr *basicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) (*entity.ClientEntity, int) {
	// Perhaps replace the memory table with a Redis Cache

	rows := client.SQLClient.Query("select ExpiryTime, Permissions from AccessTokens where AccessToken=?", token.AccessToken)
	if rows == nil {
		return nil, 1
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		expiryTime := int64(0)
		record := entity.ClientEntity{}

		err := rows.Scan(&expiryTime, &record.Permissions)
		if helper.IsError(err) {
			return nil, 2
		} else if expiryTime < time.Now().Unix() {
			return nil, 3
		}

		log.Info("Fetched client with from access token")
		return &record, 0
	}

	return nil, 2
}

func (bcr *basicClientRepository) CreateClient(clientEntity *entity.ClientEntity) *entity.ClientCredentials {
	sg := util.StringGenerator{}
	sg.Init()

	clientId := sg.GenerateString(16)
	clientSecret := sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into Clients(Id, Secret, Permissions) values(?, ?, ?)", 
		clientId, clientSecret, clientEntity.Permissions,
	)
	if res == nil {
		return nil
	}

	return &entity.ClientCredentials{
		Id: clientId, 
		Secret: clientSecret,
	}
}