package notification

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "notification_service/proto/notification"
	"notification_service/shared"
)
import "github.com/gorilla/websocket"

var Client *redis.Client

func NewNotificationController(notificationService *NotificationService) *NotificationController {
	controller := &NotificationController{NotificationService: notificationService}
	return controller
}

type NotificationController struct {
	wsConn *websocket.Conn
	UnimplementedNotificationServiceServer
	NotificationService *NotificationService
}

func (controller *NotificationController) SendNotification(ctx context.Context, notification *SendNotificationRequest) (*SendNotificationResponse, error) {
	data, err := json.Marshal(notification)
	if err != nil {
		return &SendNotificationResponse{}, err
	}
	if controller.NotificationService.CanSendNotification(notification.NotificationType, shared.StringToObjectId(notification.UserId)) {
		_, err = Client.Publish(context.Background(), "notification", data).Result()
		if err != nil {
			return &SendNotificationResponse{}, err
		}
	}
	return &SendNotificationResponse{}, nil
}

func (controller *NotificationController) CreateNotificationPreferences(ctx context.Context, req *CreateNotificationPreferencesRequest) (*CreateNotificationPreferencesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	_ = controller.NotificationService.Create(NewNotification(req))

	return &CreateNotificationPreferencesResponse{}, nil
}

func (controller *NotificationController) FindById(ctx context.Context, req *FindNotificationPreferencesByIdRequest) (*FindNotificationPreferencesByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	notification, err := controller.NotificationService.FindById(shared.StringToObjectId(req.UserId))
	if err != nil {
		return &FindNotificationPreferencesByIdResponse{}, status.Error(codes.Internal, err.Message)
	}

	return &FindNotificationPreferencesByIdResponse{UserId: notification.UserId.Hex(),
		GuestCreatedReservationRequest:     notification.GuestCreatedReservationRequest,
		GuestCanceledReservation:           notification.GuestCanceledReservation,
		GuestRatedAccommodation:            notification.GuestRatedAccommodation,
		GuestRatedHost:                     notification.GuestRatedHost,
		DistinguishedHost:                  notification.DistinguishedHost,
		HostConfirmedOrRejectedReservation: notification.HostConfirmedOrRejectedReservation}, nil
}

func (controller *NotificationController) UpdateNotificationPreferences(ctx context.Context, req *CreateNotificationPreferencesRequest) (*CreateNotificationPreferencesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	err := controller.NotificationService.Update(NewNotification(req))
	if err != nil {
		return &CreateNotificationPreferencesResponse{}, status.Error(codes.Internal, err.Message)
	}
	return &CreateNotificationPreferencesResponse{}, nil
}
