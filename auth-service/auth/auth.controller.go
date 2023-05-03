package auth

import (
	. "auth_service/auth/model"
	. "auth_service/proto/auth"
	"context"
	. "context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
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
		Id:      registered.Id.Hex(),
		Email:   registered.Email,
		Name:    registered.Name,
		Surname: registered.Surname,
		Role:    strconv.Itoa(int(registered.Role)),
	}

	return response, nil
}

func (authController *AuthController) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with user data")
	}
	if req.Token == "" {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	id, e := authController.AuthService.DecryptToken(req.Token)
	if e != nil {
		return nil, status.Error(codes.Unauthenticated, e.Message)
	}
	user, e := authController.AuthService.GetCurrentUser(id)
	if e != nil {
		return nil, status.Error(codes.NotFound, e.Message)
	}

	response := &GetUserResponse{
		Id:    user.Id.Hex(),
		Email: user.Email,
		Role:  strconv.Itoa(int(user.Role)),
	}

	return response, nil
}

func (authController *AuthController) UpdatePersonalInfo(ctx context.Context, req *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with user data")
	}
	user := UserFromUpdatePersonalInfoDto(req)
	updatedUser, e := authController.AuthService.UpdatePersonalInfo(*user)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	response := &UpdatePersonalInfoResponse{
		Id:           updatedUser.Id.Hex(),
		Email:        updatedUser.Email,
		Name:         updatedUser.Name,
		Surname:      updatedUser.Surname,
		Street:       updatedUser.Street,
		StreetNumber: updatedUser.StreetNumber,
		City:         updatedUser.City,
		ZipCode:      updatedUser.ZipCode,
		Country:      updatedUser.Country,
	}
	return response, nil

}
func (authController *AuthController) DeleteProfile(ctx context.Context, req *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	e := authController.AuthService.Delete(id)
	if e != nil {
		return &DeleteProfileResponse{Deleted: false}, status.Error(codes.Internal, e.Message)
	}
	return &DeleteProfileResponse{Deleted: true}, nil
}
