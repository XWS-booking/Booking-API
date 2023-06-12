package notification

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "notification_service/notification/model"
)

type NotificationService struct {
	NotificationRepository INotificationRepository
}

func (notificationService *NotificationService) Create(notification Notification) (primitive.ObjectID, error) {
	created, err := notificationService.NotificationRepository.Create(notification)
	if err != nil {
		return created, err
	}
	return created, nil
}

func (notificationService *NotificationService) CanSendNotification(notificationType string, id primitive.ObjectID) (bool, error) {
	notification, err := notificationService.NotificationRepository.FindById(id)
	if err != nil {
		return false, err
	}
	switch notificationType {
	case "guest_created_reservation_request":
		return notification.GuestCreatedReservationRequest, nil
	case "guest_canceled_reservation":
		return notification.GuestCanceledReservation, nil
	case "guest_rated_host":
		return notification.GuestRatedHost, nil
	case "guest_rated_accommodation":
		return notification.GuestRatedAccommodation, nil
	case "distinguished_host":
		return notification.DistinguishedHost, nil
	case "host_confirmed_or_rejected_reservation":
		return notification.HostConfirmedOrRejectedReservation, nil
	default:
		return true, nil
	}
}

func (notificationService *NotificationService) Update(notification Notification) error {
	err := notificationService.NotificationRepository.Update(notification)
	if err != nil {
		return err
	}
	return nil
}

func (notificationService *NotificationService) FindById(id primitive.ObjectID) (Notification, error) {
	notification, err := notificationService.NotificationRepository.FindById(id)
	if err != nil {
		return Notification{}, err
	}
	return notification, nil
}
