package layer

import (
	"encoding/base64"
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/util"
	"strings"

	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
)

func ClientInfoMiddleware(clientRepository repository.IClientRepository,
						res rem.IResponse,
						req rem.IRequest) *entity.ClientEntity {
	header := req.GetHeaders().Get("Authentication")
	if header == "" || len(header) < len("Basic ") {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("You must provide a Client ID and a Client Secret!"))
		return nil
	}

	encodedData := header[len("Basic "):]
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("Failed to decode Client ID and Client Secret from base64."))
		return nil
	}

	keys := strings.Split(string(decodedData), ":")

	reqObj := dto.AuthRequest{
		ClientId:     &keys[0],
		ClientSecret: &keys[1],
	}

	client := clientRepository.GetClient(reqObj.ToEntity())
	if client == nil {
		res.Status(http.StatusForbidden).JSON(util.ErrorListFromTextError("You have no permission to access this resource!"))
		return nil
	}

	return client
}

func AccessTokenMiddleware(clientRepository repository.IClientRepository,
						res rem.IResponse,
						req rem.IRequest,
						permission int64) bool {
	header := req.GetHeaders().Get("Authentication")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("You must provide an Access Token via Bearer authentication!"))
		return false
	}
	token := header[len("Bearer "):]
	clientEntity, status := clientRepository.GetClientFromAccessToken(&entity.AccessToken{AccessToken: token})

	if clientEntity == nil {
		res.Status(http.StatusUnauthorized)

		if status == 1 {
			res.JSON(util.ErrorListFromTextError("Access Token not found! Probably expired."))
		} else if status == 3 {
			res.JSON(util.ErrorListFromTextError("Access Token has expired!"))
		}

		return false
	} else if !clientEntity.CheckPermission(permission) {
		res.Status(http.StatusForbidden).JSON(util.ErrorListFromTextError("You have no permission to access this resource!"))
		return false
	}

	return true
}

// TODO: Perhaps move "Master Token" to be an Access Token with no Expiry Time (null or MAX_INT) 
//       and add a new permission to create clients
func MasterTokenMiddleware(res rem.IResponse, req rem.IRequest) bool {
	header := req.GetHeaders().Get("Authentication")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("You must provide an Access Token via Bearer authentication!"))
		return false
	}
	token := header[len("Bearer "):]

	if token != helper.Config.HTTPServer.MasterAccessToken {
		res.Status(http.StatusForbidden).JSON(util.ErrorListFromTextError("You have no permission to access this resource!"))
		return false
	}

	return true
}