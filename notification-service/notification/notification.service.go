package notification

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "notification_service/notification/model"
	"notification_service/shared"
)

type NotificationService struct {
	NotificationRepository INotificationRepository
}

func (notificationService *NotificationService) Create(notification Notification) primitive.ObjectID {
	created, error := notificationService.NotificationRepository.Create(notification)
	if error != nil {
		return created
	}
	return created
}

func (notificationService *NotificationService) CanSendNotification(notificationType string, id primitive.ObjectID) bool {
	notification, err := notificationService.NotificationRepository.FindById(id)
	if err != nil {
		return false
	}
	switch notificationType {
	case "guest_created_reservation_request":
		return notification.GuestCreatedReservationRequest
	case "guest_canceled_reservation":
		return notification.GuestCanceledReservation
	case "guest_rated_host":
		return notification.GuestRatedHost
	case "guest_rated_accommodation":
		return notification.GuestRatedAccommodation
	case "distinguished_host":
		return notification.DistinguishedHost
	case "host_confirmed_or_rejected_reservation":
		return notification.HostConfirmedOrRejectedReservation
	default:
		return true
	}
}

func (notificationService *NotificationService) Update(notification Notification) *shared.Error {
	error := notificationService.NotificationRepository.Update(notification)
	if error != nil {
		return shared.NotificationPreferencesNotUpdated()
	}
	return nil
}

func (notificationService *NotificationService) FindById(id primitive.ObjectID) (Notification, *shared.Error) {
	notification, error := notificationService.NotificationRepository.FindById(id)
	if error != nil {
		return Notification{}, shared.NotificationPreferencesNotFound()
	}
	return notification, nil
}
