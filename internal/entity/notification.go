package entity

type NotificationEntity struct {
	Id 					 int	 `json:"id"`
	TemplateID  		 int     `json:"templateId"`
	FCMRegistrationToken *string `json:"fcmRegistrationToken"`
	AppID   			 string  `json:"appId"`
	ContactType			 string  `json:"contactType"`
	ContactInfo			 string  `json:"contactInfo"`
	Title 				 string  `json:"title"`
	Message 		   	 string  `json:"message"`
	SentTime			 int     `json:"sentTime"`
}