package impl

import (
	log "github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"time"
)

type BasicClientRepository struct {
	sg *util.StringGenerator
}

func (bcr *BasicClientRepository) Init() {
	bcr.sg = util.NewStringGenerator()
}

func (bcr *BasicClientRepository) GetClient(credentials *entity.ClientCredentials) *entity.ClientEntity {
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

func (bcr *BasicClientRepository) UpdateClient(clientID *string, clientEntity *entity.ClientEntity) int {
	res := client.SQLClient.Exec(
			"update Clients set Permissions=? where Id=?",
			clientEntity.Permissions, *clientID)
	if res == nil {
		return 2
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 2
	} else if affectedRows <= 0 {
		return 1
	}

	return 0
}

func (bcr *BasicClientRepository) DeleteClient(clientID *string) int {
	res := client.SQLClient.Exec("delete from Clients where Id=?", *clientID)
	if res == nil {
		return 2
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 2
	} else if affectedRows <= 0 {
		return 1
	}

	return 0
}

func (bcr *BasicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) *entity.AccessToken {
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

func (bcr *BasicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) (*entity.ClientEntity, int) {
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

func (bcr *BasicClientRepository) CreateClient(clientEntity *entity.ClientEntity) *entity.ClientCredentials {
	sg := util.NewStringGenerator()

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