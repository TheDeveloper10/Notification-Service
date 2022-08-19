package repository

import (
	log "github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
)

type ClientRepository interface {
	GetClient(*entity.ClientCredentials) 	  *entity.ClientEntity
	GenerateAccessToken(*entity.ClientEntity) string
}

type basicClientRepository struct{}


func NewClientRepository() ClientRepository {
	return &basicClientRepository{}
}

func (bcr *basicClientRepository) GetClient(credentials *entity.ClientCredentials) *entity.ClientEntity {
	stmt, err := client.SQLClient.Prepare("select * from Clients where Id=? and Secret=?")
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
		err3 := rows.Scan(&record.Id, &record.Secret, &record.Permissions, &record.CreationTime)
		if helper.IsError(err3) {
			return nil
		}

		log.Info("Fetched client with id " + credentials.Id)
		return &record
	}

	return nil
}

func (bcr *basicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) string {

	return ""
}