package notification

import (
	. "notification_service/notification/model"
	notificationGrpc "notification_service/proto/notification"
	"notification_service/shared"
)

func NewNotification(req *notificationGrpc.CreateNotificationPreferencesRequest) Notification {
	return Notification{
		UserId:                             shared.StringToObjectId(req.UserId),
		GuestCreatedReservationRequest:     req.GuestCreatedReservationRequest,
		GuestCanceledReservation:           req.GuestCanceledReservation,
		GuestRatedHost:                     req.GuestRatedHost,
		GuestRatedAccommodation:            req.GuestRatedAccommodation,
		DistinguishedHost:                  req.DistinguishedHost,
		HostConfirmedOrRejectedReservation: req.HostConfirmedOrRejectedReservation,
	}
}
