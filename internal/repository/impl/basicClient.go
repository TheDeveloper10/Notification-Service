package impl

import (
	log "github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/data/entity"
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

func (bcr *BasicClientRepository) GetClient(credentials *entity.ClientCredentials) (*entity.ClientEntity, util.RepoStatusCode) {
	rows := client.SQLClient.Query(
		"select Permissions from Clients where Id=? and Secret=?",
		credentials.Id, credentials.Secret,
	)
	if rows == nil {
		return nil, util.RepoStatusError
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		record := entity.ClientEntity{}
		err3 := rows.Scan(&record.Permissions)
		if helper.IsError(err3) {
			return nil, util.RepoStatusError
		}

		log.Info("Fetched client with id " + credentials.Id)
		return &record, util.RepoStatusSuccess
	}

	return nil, util.RepoStatusNotFound
}

func (bcr *BasicClientRepository) UpdateClient(clientID *string, clientEntity *entity.ClientEntity) util.RepoStatusCode {
	res := client.SQLClient.Exec(
			"update Clients set Permissions=? where Id=?",
			clientEntity.Permissions, *clientID)
	if res == nil {
		return util.RepoStatusError
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return util.RepoStatusError
	} else if affectedRows <= 0 {
		return util.RepoStatusNotFound
	}

	return util.RepoStatusSuccess
}

func (bcr *BasicClientRepository) DeleteClient(clientID *string) util.RepoStatusCode {
	res := client.SQLClient.Exec("delete from Clients where Id=?", *clientID)
	if res == nil {
		return util.RepoStatusError
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return util.RepoStatusError
	} else if affectedRows <= 0 {
		return util.RepoStatusNotFound
	}

	return util.RepoStatusSuccess
}

func (bcr *BasicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) (*entity.AccessToken, util.RepoStatusCode) {
	token := bcr.sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into AccessTokens(AccessToken, Permissions, ExpiryTime) values(?, ?, unix_timestamp() + ?)",
		token, clientEntity.Permissions, helper.Config.HTTPServer.AccessTokenExpiryTime,
	)
	if res == nil {
		return nil, util.RepoStatusError
	}

	log.Info("Generated a new access token")
	return &entity.AccessToken{
		AccessToken: token,
	}, util.RepoStatusSuccess
}

func (bcr *BasicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) (*entity.ClientEntity, util.RepoStatusCode) {
	// Perhaps replace the memory table with a Redis Cache

	rows := client.SQLClient.Query("select ExpiryTime, Permissions from AccessTokens where AccessToken=?", token.AccessToken)
	if rows == nil {
		return nil, util.RepoStatusError
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		expiryTime := int64(0)
		record := entity.ClientEntity{}

		err := rows.Scan(&expiryTime, &record.Permissions)
		if helper.IsError(err) {
			return nil, util.RepoStatusError
		} else if expiryTime < time.Now().Unix() {
			return nil, util.RepoStatusExpired
		}

		log.Info("Fetched client with from access token")
		return &record, util.RepoStatusSuccess
	}

	return nil, util.RepoStatusNotFound
}

func (bcr *BasicClientRepository) CreateClient(clientEntity *entity.ClientEntity) (*entity.ClientCredentials, util.RepoStatusCode) {
	sg := util.NewStringGenerator()

	clientId := sg.GenerateString(16)
	clientSecret := sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into Clients(Id, Secret, Permissions) values(?, ?, ?)",
		clientId, clientSecret, clientEntity.Permissions,
	)
	if res == nil {
		return nil, util.RepoStatusError
	}

	return &entity.ClientCredentials{
		Id: clientId,
		Secret: clientSecret,
	}, util.RepoStatusSuccess
}