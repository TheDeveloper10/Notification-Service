package entity

import "github.com/golang-jwt/jwt"

type AuthTokens struct {
	AccessToken    string
	RefresherToken string
}

type ClientCredentials struct {
	Id     string
	Secret string
}

type ClientEntity struct {
	ClientId    string `json:"cid"`
	Permissions int64  `json:"pm"`
}

type ClientEntityClaims struct {
	ClientEntity
	jwt.StandardClaims
}

func (ce *ClientEntity) CheckPermission(permission int64) bool {
	return ce.Permissions & permission > 0
}


const (
	PermissionSendNotifications     	= 1
	PermissionSendNotificationsKey  	= "send_notifications"
	PermissionReadSentNotifications 	= 2
	PermissionReadSentNotificationsKey 	= "read_sent_notifications"
	PermissionCreateTemplates       	= 4
	PermissionCreateTemplatesKey       	= "create_templates"
	PermissionReadTemplates         	= 8
	PermissionReadTemplatesKey         	= "read_templates"
	PermissionUpdateTemplates       	= 16
	PermissionUpdateTemplatesKey       	= "update_templates"
	PermissionDeleteTemplates       	= 32
	PermissionDeleteTemplatesKey       	= "delete_templates"
	PermissionAll                       = 63
)

func PermissionKeyToInt(key string) int {
	switch key {
	case PermissionSendNotificationsKey:
		return PermissionSendNotifications
	case PermissionReadSentNotificationsKey:
		return PermissionReadSentNotifications
	case PermissionCreateTemplatesKey:
		return PermissionCreateTemplates
	case PermissionReadTemplatesKey:
		return PermissionReadTemplates
	case PermissionUpdateTemplatesKey:
		return PermissionUpdateTemplates
	case PermissionDeleteTemplatesKey:
		return PermissionDeleteTemplates
	default:
		return -1
	}
}