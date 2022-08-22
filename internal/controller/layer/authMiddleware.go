package layer

import (
	"net/http"
	"strings"

	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/util/iface"
)

func ClientInfoMiddleware(clientRepository repository.ClientRepository,
						res iface.IResponseWriter,
						req *http.Request) *entity.ClientEntity {
	header := req.Header.Get("Authentication")
	if header == "" || len(header) < len("Basic ") {
		res.Status(http.StatusUnauthorized).TextError("You must provide a Client ID and a Client Secret!")
		return nil
	}
	keys := strings.Split(header[len("Basic "):], ":")

	reqObj := dto.AuthRequest{
		ClientId:     &keys[0],
		ClientSecret: &keys[1],
	}

	client := clientRepository.GetClient(reqObj.ToEntity())
	if client == nil {
		res.Status(http.StatusForbidden).TextError("You have no permission to access this resource!")
		return nil
	}

	return client
}

func AccessTokenMiddleware(clientRepository repository.ClientRepository,
						res iface.IResponseWriter,
						req *http.Request,
						permission int) bool {
	header := req.Header.Get("Authentication")
	if header == "" || len(header) <= len("Bearer ") {
		res.Status(http.StatusUnauthorized).TextError("You must provide an Access Token!")
		return false
	}
	token := header[len("Bearer "):]
	clientEntity, status := clientRepository.GetClientFromAccessToken(&entity.AccessToken{AccessToken: token})

	if clientEntity == nil {
		res.Status(http.StatusUnauthorized)

		if status == 1 {
			res.TextError("Access Token not found! Probably expired.")
		} else if status == 3 {
			res.TextError("Access Token has expired!")
		}

		return false
	} else if !clientEntity.CheckPermission(permission) {
		res.Status(http.StatusForbidden).TextError("You have no permission to access this resource!")
		return false
	}

	return true
}

// TODO: Perhaps move "Master Token" to be an Access Token with no Expiry Time (null or MAX_INT) 
//       and add a new permission to create clients
func MasterTokenMiddleware(res iface.IResponseWriter, req *http.Request) bool {
	header := req.Header.Get("Authentication")
	if header == "" || len(header) <= len("Bearer ") {
		res.Status(http.StatusUnauthorized).TextError("You must provide an Access Token!")
		return false
	}
	token := header[len("Bearer "):]

	if token != helper.Config.HTTPServer.MasterAccessToken {
		res.Status(http.StatusForbidden).TextError("You have no permission to access this resource!")
		return false
	}

	return true
}