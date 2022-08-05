package entity

type NotificationEntity struct {
	Id 			int
	TemplateID  int
	UserID  	string
	AppID   	string
	ContactType string
	ContactInfo string
	Title		string
	Message 	string
	SentTime   	int
}