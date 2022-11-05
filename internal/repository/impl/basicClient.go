package impl

import (
	"github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"time"
)

type BasicClientRepository struct {
	sg *util.StringGenerator
}

func (bcr *BasicClientRepository) Init() {
	bcr.sg = util.NewStringGenerator()
}

func (bcr *BasicClientRepository) GetClient(credentials *entity.ClientCredentials) (*entity.ClientEntity, code.StatusCode) {
	rows := client.SQLClient.Query(
		"select Permissions from Clients where Id=? and Secret=?",
		credentials.Id, credentials.Secret,
	)
	if rows == nil {
		return nil, code.StatusError
	}
	defer util.HandledClose(rows)

	if rows.Next() {
		record := entity.ClientEntity{}
		err3 := rows.Scan(&record.Permissions)
		if util.ManageError(err3) {
			return nil, code.StatusError
		}

		logrus.Info("Fetched client with id " + credentials.Id)
		return &record, code.StatusSuccess
	}

	return nil, code.StatusNotFound
}

func (bcr *BasicClientRepository) UpdateClient(clientID *string, clientEntity *entity.ClientEntity) code.StatusCode {
	res := client.SQLClient.Exec(
			"update Clients set Permissions=? where Id=?",
			clientEntity.Permissions, *clientID)
	if res == nil {
		return code.StatusError
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return code.StatusError
	} else if affectedRows <= 0 {
		return code.StatusNotFound
	}

	return code.StatusSuccess
}

func (bcr *BasicClientRepository) DeleteClient(clientID *string) code.StatusCode {
	res := client.SQLClient.Exec("delete from Clients where Id=?", *clientID)
	if res == nil {
		return code.StatusError
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return code.StatusError
	} else if affectedRows <= 0 {
		return code.StatusNotFound
	}

	return code.StatusSuccess
}

func (bcr *BasicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) (*entity.AccessToken, code.StatusCode) {
	token := bcr.sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into AccessTokens(AccessToken, Permissions, ExpiryTime) values(?, ?, unix_timestamp() + ?)",
		token, clientEntity.Permissions, util.Config.HTTPServer.AccessTokenExpiryTime,
	)
	if res == nil {
		return nil, code.StatusError
	}

	logrus.Info("Generated a new access token")
	return &entity.AccessToken{
		AccessToken: token,
	}, code.StatusSuccess
}

func (bcr *BasicClientRepository) GetClientFromAccessToken(token *entity.AccessToken) (*entity.ClientEntity, code.StatusCode) {
	// Perhaps replace the memory table with a Redis Cache

	rows := client.SQLClient.Query("select ExpiryTime, Permissions from AccessTokens where AccessToken=?", token.AccessToken)
	if rows == nil {
		return nil, code.StatusError
	}
	defer util.HandledClose(rows)

	if rows.Next() {
		expiryTime := int64(0)
		record := entity.ClientEntity{}

		err := rows.Scan(&expiryTime, &record.Permissions)
		if util.ManageError(err) {
			return nil, code.StatusError
		} else if expiryTime < time.Now().Unix() {
			return nil, code.StatusExpired
		}

		logrus.Info("Fetched client with from access token")
		return &record, code.StatusSuccess
	}

	return nil, code.StatusNotFound
}

func (bcr *BasicClientRepository) CreateClient(clientEntity *entity.ClientEntity) (*entity.ClientCredentials, code.StatusCode) {
	sg := util.NewStringGenerator()

	clientId := sg.GenerateString(16)
	clientSecret := sg.GenerateString(128)

	res := client.SQLClient.Exec(
		"insert into Clients(Id, Secret, Permissions) values(?, ?, ?)",
		clientId, clientSecret, clientEntity.Permissions,
	)
	if res == nil {
		return nil, code.StatusError
	}

	return &entity.ClientCredentials{
		Id: clientId,
		Secret: clientSecret,
	}, code.StatusSuccess
}