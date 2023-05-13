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

type UpdatePersonalInfoDto struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	City         string `json:"city"`
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
	Username     string `json:"username"`
}

type UpdatePersonalInfoHandler struct {
	authClientAddress string
}

func NewUpdatePersonalInfoHandler(authClientAddress string) Handler {
	return &UpdatePersonalInfoHandler{
		authClientAddress: authClientAddress,
	}
}

func (handler *UpdatePersonalInfoHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/auth/personal/update", TokenValidationMiddleware(RolesMiddleware([]UserRole{0, 1}, UserMiddleware(handler.UpdatePersonalInfo))))
	if err != nil {
		panic(err)
	}
}

func (handler *UpdatePersonalInfoHandler) UpdatePersonalInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	var body UpdatePersonalInfoDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	req := &gateway.UpdatePersonalInfoRequest{
		Id:           body.Id,
		Name:         body.Name,
		Surname:      body.Surname,
		Email:        body.Email,
		Password:     "",
		Street:       body.Street,
		StreetNumber: body.StreetNumber,
		City:         body.City,
		ZipCode:      body.ZipCode,
		Country:      body.Country,
		Username:     body.Username,
	}
	fmt.Println(req.Id)

	res, err := authClient.UpdatePersonalInfo(context.TODO(), req)
	if err != nil {
		shared.BadRequest(w, "Error when updating personal info!")
		return
	}
	shared.Ok(&w, res)
}
