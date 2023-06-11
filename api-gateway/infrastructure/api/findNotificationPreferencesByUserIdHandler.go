package api

import (
	"context"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	ctx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type FindNotificationPreferencesByUserIdHandler struct {
	notificationClientAddress string
}

func NewFindNotificationPreferencesByUserHandler(notificationClientAddress string) Handler {
	return &FindNotificationPreferencesByUserIdHandler{
		notificationClientAddress: notificationClientAddress,
	}
}

func (handler *FindNotificationPreferencesByUserIdHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/notification/preferences", TokenValidationMiddleware(RolesMiddleware([]UserRole{0, 1}, UserMiddleware(handler.Find))))
	if err != nil {
		panic(err)
	}
}

func (handler *FindNotificationPreferencesByUserIdHandler) Find(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := ctx.Get(r, "id").(string)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	notificationPreferences, err := notificationClient.FindById(context.TODO(), &gateway.FindNotificationPreferencesByIdRequest{UserId: id})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, notificationPreferences)
}
