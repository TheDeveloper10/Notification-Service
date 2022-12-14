package layer

import (
	"encoding/base64"
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"strings"

	"notification-service/internal/repository"
)

func ClientInfoMiddleware(clientRepository repository.IClientRepository,
						res rem.IResponse,
						req rem.IRequest) *entity.ClientEntity {
	header := req.GetHeaders().Get("Authorization")
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
		ClientId:     keys[0],
		ClientSecret: keys[1],
	}

	client, status := clientRepository.GetClient(reqObj.ToEntity())
	if status == code.StatusSuccess {
		return client
	} else if status == code.StatusNotFound {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Client not found!"))
	} else if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
	}

	return nil
}

func AccessTokenMiddleware(clientRepository repository.IClientRepository,
						res rem.IResponse,
						req rem.IRequest,
						permission int64) bool {
	header := req.GetHeaders().Get("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("You must provide an Access Token via Bearer authorization!"))
		return false
	}
	token := header[len("Bearer "):]
	clientEntity, status := clientRepository.ExtractClientFromToken(&token, &util.Config.Service.AccessTokenSecret)

	if clientEntity == nil || status != code.StatusSuccess {
		res.Status(http.StatusUnauthorized)

		if status == code.StatusNotFound {
			res.JSON(util.ErrorListFromTextError("Access Token not found! Probably expired."))
		} else if status == code.StatusError {
			res.JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
		} else if status == code.StatusExpired {
			res.JSON(util.ErrorListFromTextError("Access Token has expired!"))
		}

		return false
	} else if !clientEntity.CheckPermission(permission) {
		res.Status(http.StatusForbidden).JSON(util.ErrorListFromTextError("You have no permission to access this resource!"))
		return false
	}

	return true
}

func MasterTokenMiddleware(res rem.IResponse, req rem.IRequest) bool {
	header := req.GetHeaders().Get("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		res.Status(http.StatusUnauthorized).JSON(util.ErrorListFromTextError("You must provide an Access Token via Bearer authorization!"))
		return false
	}
	token := header[len("Bearer "):]

	if token != util.Config.HTTPServer.MasterAccessToken {
		res.Status(http.StatusForbidden).JSON(util.ErrorListFromTextError("You have no permission to access this resource!"))
		return false
	}

	return true
}