package api

import (
	"context"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UpdateNotificationPreferencesHandler struct {
	notificationClientAddress string
}

func NewUpdateNotificationPreferencesHandler(notificationClientAddress string) Handler {
	return &UpdateNotificationPreferencesHandler{
		notificationClientAddress: notificationClientAddress,
	}
}

func (handler *UpdateNotificationPreferencesHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PUT", "/api/notification/preferences", TokenValidationMiddleware(RolesMiddleware([]UserRole{0, 1}, UserMiddleware(handler.Update))))
	if err != nil {
		panic(err)
	}
}

func (handler *UpdateNotificationPreferencesHandler) Update(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body *gateway.CreateNotificationPreferencesRequest
	err := DecodeBody(r, &body)
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err = notificationClient.UpdateNotificationPreferences(context.TODO(), body)
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, "")
}
