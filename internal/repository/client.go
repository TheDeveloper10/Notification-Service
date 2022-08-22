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
		sg: &util.StringGenerator{},
	}
}

func (bcr *basicClientRepository) GetClient(credentials *entity.ClientCredentials) *entity.ClientEntity {
	stmt, err := client.SQLClient.Prepare("select Permissions from Clients where Id=? and Secret=?")
	if helper.IsError(err) {
		return nil
	}

	rows, err2 := stmt.Query(credentials.Id, credentials.Secret)
	if helper.IsError(err2) {
		return nil
	}
	defer helper.HandledClose(rows)

	for rows.Next() {
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
	token := bcr.sg.GenerateString(helper.Config.HTTPServer.AccessTokenKeyLen)

	stmt, err1 := client.SQLClient.Prepare("insert into AccessTokens(AccessToken, Permissions) values(?, ?)")
	if helper.IsError(err1) {
		return nil
	}
	defer helper.HandledClose(stmt)

	_, err2 := stmt.Exec(token, clientEntity.Permissions)
	if helper.IsError(err2) {
		return nil
	}

	return &entity.AccessToken{
		AccessToken: token,
	}
}

func (bcr *basicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) *entity.ClientEntity {
	// Perhaps replace the memory table with a Redis Cache

	stmt, err := client.SQLClient.Prepare("select Permissions from AccessTokens where AccessToken=?")
	if helper.IsError(err) {
		return nil
	}

	rows, err2 := stmt.Query(token.AccessToken)
	if helper.IsError(err2) {
		return nil
	}
	defer helper.HandledClose(rows)

	for rows.Next() {
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
