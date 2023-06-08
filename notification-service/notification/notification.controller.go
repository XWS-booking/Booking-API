package notification

import (
	"context"
	"net/http"
	. "notification_service/proto/notification"
)
import "github.com/gorilla/websocket"

var Client *websocket.Conn = &websocket.Conn{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		return true
	},
}

func NewNotificationController(notificationService *NotificationService) *NotificationController {
	controller := &NotificationController{NotificationService: notificationService}
	return controller
}

type NotificationController struct {
	wsConn *websocket.Conn
	UnimplementedNotificationServiceServer
	NotificationService *NotificationService
}

func (controller *NotificationController) SendNotification(ctx context.Context, notification *NotificationRequest) (*NotificationResponse, error) {
	err := Client.WriteJSON(notification)
	if err != nil {
		return &NotificationResponse{Status: "lose"}, err
	}

	return &NotificationResponse{Status: "ok"}, nil
}
