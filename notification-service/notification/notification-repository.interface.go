package notification

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "notification_service/notification/model"
)

type INotificationRepository interface {
	Create(notification Notification) (primitive.ObjectID, error)
	FindById(id primitive.ObjectID) (Notification, error)
	Update(notification Notification) error
}
