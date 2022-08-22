package layer

import (
	"net/http"

	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util/iface"
)

func AuthMiddleware(clientRepository repository.ClientRepository,
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
		if status == 1 {
			res.Status(http.StatusUnauthorized).TextError("Access Token not found! Probably expired!")
		} else if status == 3 {
			res.Status(http.StatusUnauthorized).TextError("Access Token has expired!")
		}

		res.Status(http.StatusUnauthorized)
		return false
	} else if !clientEntity.CheckPermission(permission) {
		res.Status(http.StatusForbidden).TextError("You have no permission to access this resource!")
		return false
	}

	return true
}
