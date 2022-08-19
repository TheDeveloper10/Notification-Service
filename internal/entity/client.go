package entity

type ClientCredentials struct {
	Id     string
	Secret string
}

type ClientEntity struct {
	ClientCredentials
	Permissions  int
	CreationTime int
}

const (
	PermissionSendNotifications     = 1
	PermissionReadSentNotifications = 2
	PermissionCreateTemplates       = 4
	PermissionUpdateTemplates       = 8
	PermissionDeleteTemplates       = 16
)

func (ce *ClientEntity) CheckPermission(permission int) bool {
	return ce.Permissions & permission > 0
}