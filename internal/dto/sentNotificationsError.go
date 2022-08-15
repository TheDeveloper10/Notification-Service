package dto

type SentNotificationsError struct {
	SentNotifications int    `json:"sentNotifications"`
	Error             string `json:"error"`
}
