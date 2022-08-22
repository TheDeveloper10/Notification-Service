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
		res.Status(http.StatusUnauthorized)
		return false
	}
	token := header[len("Bearer "):]
	clientEntity := clientRepository.GetClientFromAccessToken(&entity.AccessToken{AccessToken: token})

	if !clientEntity.CheckPermission(permission) {
		res.Status(http.StatusForbidden)
		return false
	}

	return true
}
