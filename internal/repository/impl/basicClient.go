package impl

import (
	"fmt"
	"github.com/golang-jwt/jwt"
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
		record := entity.ClientEntity{
			ClientId: credentials.Id,
		}
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




func (bcr *BasicClientRepository) VerifyToken(token *string, secret *string) code.StatusCode {
	_, status := bcr.extractMapFromToken(token, secret)
	return status
}

func (bcr *BasicClientRepository) GenerateToken(clientEntity *entity.ClientEntity, secret *string, expiry int) (*string, code.StatusCode) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.ClientEntityClaims{
		ClientEntity: *clientEntity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Second).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(*secret))
	if util.ManageError(err) {
		return nil, code.StatusError
	}

	return &tokenString, code.StatusSuccess
}

func (bcr *BasicClientRepository) ExtractClientFromToken(token *string, secret *string) (*entity.ClientEntity, code.StatusCode) {
	claims, status := bcr.extractMapFromToken(token, secret)
	if status != code.StatusSuccess {
		return nil, status
	}

	return &claims.ClientEntity, code.StatusSuccess
}


func (bcr *BasicClientRepository) extractMapFromToken(target *string, secret *string) (*entity.ClientEntityClaims, code.StatusCode) {
	token, err := jwt.ParseWithClaims(*target, &entity.ClientEntityClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(*secret), nil
	})

	if util.ManageError(err) {
		return nil, code.StatusError
	}

	claims, ok := token.Claims.(*entity.ClientEntityClaims)
	if ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, code.StatusExpired
		}

		return claims, code.StatusSuccess
	} else {
		return nil, code.StatusError
	}
}
