package notification

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	. "notification_service/opentelementry"
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
	_, span := Tp.Tracer(ServiceName).Start(ctx, "sendNotificion")
	defer func() { span.End() }()
	data, err := json.Marshal(notification)
	if err != nil {
		HttpError(err, span, http.StatusBadRequest)
		return &SendNotificationResponse{}, status.Error(http.StatusBadRequest, err.Error())
	}
	canSend, err := controller.NotificationService.CanSendNotification(notification.NotificationType, shared.StringToObjectId(notification.UserId))
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return &SendNotificationResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}
	if canSend {
		_, err = Client.Publish(context.Background(), "notification", data).Result()
		if err != nil {
			HttpError(err, span, http.StatusInternalServerError)
			return &SendNotificationResponse{}, status.Error(http.StatusInternalServerError, err.Error())
		}
	}
	return &SendNotificationResponse{}, nil
}

func (controller *NotificationController) CreateNotificationPreferences(ctx context.Context, req *CreateNotificationPreferencesRequest) (*CreateNotificationPreferencesResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "createNotification")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	_, err := controller.NotificationService.Create(NewNotification(req))
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return &CreateNotificationPreferencesResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}
	return &CreateNotificationPreferencesResponse{}, nil
}

func (controller *NotificationController) FindById(ctx context.Context, req *FindNotificationPreferencesByIdRequest) (*FindNotificationPreferencesByIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findById")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	notification, err := controller.NotificationService.FindById(shared.StringToObjectId(req.UserId))
	if err != nil {
		status.Error(http.StatusInternalServerError, err.Error())
		return &FindNotificationPreferencesByIdResponse{}, status.Error(http.StatusInternalServerError, err.Error())
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
	_, span := Tp.Tracer(ServiceName).Start(ctx, "updateNotificationPreferences")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	err := controller.NotificationService.Update(NewNotification(req))
	if err != nil {
		status.Error(http.StatusInternalServerError, err.Error())
		return &CreateNotificationPreferencesResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}
	return &CreateNotificationPreferencesResponse{}, nil
}
