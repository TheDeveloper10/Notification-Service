package dto

type SentNotificationsError struct {
	SentNotifications int    `json:"sentNotifications"`
	Error1            string `json:"error-1"`
	Error2            string `json:"error-2"`
}
