package auth

import (
	. "auth_service/auth/model"
	. "auth_service/proto/auth"
	. "auth_service/shared"
)

func UserFromRegistrationDto(dto *RegistrationRequest) *User {
	role := GUEST
	if dto.Role == 1 {
		role = HOST
	}
	return &User{
		Email:        dto.Email,
		Name:         dto.Name,
		Surname:      dto.Surname,
		Password:     dto.Password,
		Username:     dto.Username,
		Street:       dto.Street,
		StreetNumber: dto.StreetNumber,
		City:         dto.City,
		ZipCode:      dto.ZipCode,
		Country:      dto.Country,
		Role:         role,
	}
}

func UserFromSignInDto(dto *SignInRequest) *User {
	return &User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func UserFromUpdatePersonalInfoDto(dto *UpdatePersonalInfoRequest) *User {
	return &User{
		Id:           StringToObjectId(dto.Id),
		Name:         dto.Name,
		Surname:      dto.Surname,
		Password:     dto.Password,
		Email:        dto.Email,
		Street:       dto.Street,
		StreetNumber: dto.StreetNumber,
		City:         dto.City,
		ZipCode:      dto.ZipCode,
		Country:      dto.Country,
		Username:     dto.Username,
	}
}
