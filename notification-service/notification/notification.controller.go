package notification

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"net/http"
	. "notification_service/proto/notification"
)
import "github.com/gorilla/websocket"

var Client *redis.Client

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
	data, err := json.Marshal(notification)
	if err != nil {
		return &NotificationResponse{}, err
	}
	return &NotificationResponse{Status: "ok"}, Client.Publish(context.Background(), "notification", data).Err()
}
