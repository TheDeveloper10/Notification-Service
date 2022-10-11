package dto

type SendNotificationErrorData struct {
	TargetId int	  `json:"targetId"`
	Messages []string `json:"messages"`
}

type SendNotificationsError struct {
	Errors 						  []SendNotificationErrorData `json:"errors"`
	SuccessfullySentNotifications int                         `json:"successfullySentNotifications"`
	FailedNotifications 		  int 					  	  `json:"failedNotifications"`
}
