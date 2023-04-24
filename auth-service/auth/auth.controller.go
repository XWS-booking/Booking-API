package auth

import (
	. "auth_service/auth/model"
	. "auth_service/proto/auth"
	"context"
	. "context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateAuthController(authService *AuthService) *AuthController {
	controller := &AuthController{AuthService: authService}
	return controller
}

type AuthController struct {
	UnimplementedAuthServiceServer
	AuthService *AuthService
}

func (authController *AuthController) SignIn(ctx Context, req *SignInRequest) (*SignInResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with user data")
	}

	user := UserFromSignInDto(req)
	token, e := authController.AuthService.SignIn(user.Email, user.Password)

	if e != nil {
		return nil, status.Error(codes.Unauthenticated, e.Message)
	}

	fmt.Println(token)
	response := &SignInResponse{
		AccessToken: token,
	}

	return response, nil
}

func (authController *AuthController) Register(ctx context.Context, req *RegistrationRequest) (*RegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with user data")
	}

	user := UserFromRegistrationDto(req)
	user.Role = HOST
	registered, e := authController.AuthService.Register(*user)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	response := &RegistrationResponse{
		Id:      registered.Id.String(),
		Email:   registered.Email,
		Name:    registered.Name,
		Surname: registered.Surname,
		Role:    string(registered.Role),
	}

	return response, nil
}
