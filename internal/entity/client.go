package entity

type ClientCredentials struct {
	Id     string
	Secret string
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type ClientEntity struct {
	Permissions int
}

const (
	PermissionSendNotifications     = 1
	PermissionReadSentNotifications = 2
	PermissionCreateTemplates       = 4
	PermissionReadTemplates         = 8
	PermissionUpdateTemplates       = 16
	PermissionDeleteTemplates       = 32
)

func (ce *ClientEntity) CheckPermission(permission int) bool {
	return ce.Permissions & permission > 0
}
