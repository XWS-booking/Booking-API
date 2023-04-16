package auth

import (
	. "auth_service/auth/model"
	. "auth_service/proto/auth"
	authGrpc "auth_service/proto/auth"
	"context"
	. "context"

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
	userProto := req.User
	if userProto == nil {
		return nil, status.Error(codes.Aborted, "Something wront with user data")
	}
	var user User
	user.MapFromProto(userProto)

	token, e := authController.AuthService.SignIn(user.Email, user.Password)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}
	var tokenResp Token
	tokenResp.AccessToken = token
	response := &SignInResponse{
		AccessToken: &tokenResp,
	}
	return response, nil
}

func (authController *AuthController) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	userProto := req.User
	if userProto == nil {
		return nil, status.Error(codes.Aborted, "Something wront with user data")
	}
	var user User
	user.MapFromProto(userProto)
	user, e := authController.AuthService.Register(user)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	var userResp authGrpc.UserDto
	userResp.Id = string(user.Id.Hex())
	userResp.Name = user.Name
	userResp.Surname = user.Surname
	userResp.Email = user.Email
	userResp.Role = authGrpc.Role(user.Role)

	response := &RegisterResponse{
		UserDto: &userResp,
	}
	return response, nil
}
