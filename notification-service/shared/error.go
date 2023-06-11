package shared

type Error struct {
	Message string
}

func NotificationPreferencesNotUpdated() *Error {
	return &Error{Message: "Notification preferences can't be updated!"}
}

func NotificationPreferencesNotFound() *Error {
	return &Error{Message: "Notification preferences not found!"}
}
