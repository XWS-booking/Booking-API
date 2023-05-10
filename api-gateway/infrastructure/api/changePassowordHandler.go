package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type ChangePasswordDto struct {
	Id          string `json:"id"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordHandler struct {
	authClientAddress string
}

func NewChangePasswordHandler(authClientAddress string) Handler {
	return &ChangePasswordHandler{
		authClientAddress: authClientAddress,
	}
}

func (handler *ChangePasswordHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/auth/password/change", TokenValidationMiddleware(RolesMiddleware([]UserRole{0, 1}, UserMiddleware(handler.ChangePassword))))
	if err != nil {
		panic(err)
	}
}

func (handler *ChangePasswordHandler) ChangePassword(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	var body ChangePasswordDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	req := &gateway.ChangePasswordRequest{
		Id:          body.Id,
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
	}

	res, err := authClient.ChangePassword(context.TODO(), req)
	if err != nil {
		shared.BadRequest(w, "Error when updating password!")
		return
	}
	shared.Ok(&w, res)
}
