package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type LoginHandler struct {
	authClientAddress string
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginHandler(authClientAddress string) Handler {
	return &LoginHandler{
		authClientAddress: authClientAddress,
	}
}

func (handler *LoginHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/auth/signin", handler.Login)
	if err != nil {
		panic(err)
	}
}

func (handler *LoginHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	loginReq := &gateway.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, e := authClient.SignIn(context.TODO(), loginReq)
	if e != nil {
		panic(e)
	}
	shared.Ok(&w, res)
}
