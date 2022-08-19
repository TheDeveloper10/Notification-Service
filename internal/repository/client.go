package repository

import (
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"time"
)

type ClientRepository interface {
	GetClient(*entity.ClientCredentials) 	  *entity.ClientEntity
	GenerateAccessToken(*entity.ClientEntity) *entity.AccessToken
	ValidateAccessToken(*entity.AccessToken)  bool
}

type basicClientRepository struct{
	keyManager 	  util.KeyManager
	cryptoManager util.CryptoManager
}


func NewClientRepository() ClientRepository {
	repo := basicClientRepository{
		keyManager:    util.KeyManager{},
		cryptoManager: util.CryptoManager{},
	}
	repo.init()
	return &repo
}

const (
	JWTSigningKey 	 = "jwt_signing_key"
	JWTEncryptionKey = "jwt_enc_key"
)

func (bcr *basicClientRepository) init() {
	bcr.keyManager.Init()
	bcr.keyManager.GenerateKey(JWTSigningKey, 	 helper.Config.HTTPServer.AccessTokenKeyLen)
	bcr.keyManager.GenerateKey(JWTEncryptionKey, helper.Config.HTTPServer.AccessTokenEncryptionKeyLen)
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

func (bcr *basicClientRepository) GenerateAccessToken(clientEntity *entity.ClientEntity) *entity.AccessToken {
	expirationTime := time.Now().Add(time.Minute * time.Duration(helper.Config.HTTPServer.AccessTokenExpiryTime))

	claims := entity.ClientClaims{
		ClientId: 	 clientEntity.Id,
		Permissions: clientEntity.Permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(*bcr.keyManager.GetKey(JWTSigningKey)))
	if helper.IsError(err) {
		return nil
	}

	encryptedToken := bcr.cryptoManager.Encrypt(&tokenString, bcr.keyManager.GetKey(JWTEncryptionKey))
	if encryptedToken != nil {
		log.Info("Generated a new access token")
	}
	return &entity.AccessToken{
		AccessToken: *encryptedToken,
	}
}

func (bcr *basicClientRepository) ValidateAccessToken(token *entity.AccessToken) bool {
	jwtToken := bcr.cryptoManager.Decrypt(&token.AccessToken, bcr.keyManager.GetKey(JWTEncryptionKey))
	if jwtToken == nil {
		return false
	}
	claims := entity.ClientClaims{}

	resToken, err := jwt.ParseWithClaims(*jwtToken, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return bcr.keyManager.GetKey(JWTSigningKey), nil
	})

	return !helper.IsError(err) && resToken.Valid
}