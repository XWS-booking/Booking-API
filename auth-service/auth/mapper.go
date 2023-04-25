package auth

import (
	. "auth_service/auth/model"
	. "auth_service/proto/auth"
)

func UserFromRegistrationDto(dto *RegistrationRequest) *User {
	return &User{
		Email:    dto.Email,
		Name:     dto.Name,
		Surname:  dto.Surname,
		Password: dto.Password,
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
		Id: dto.Id,
		Name: dto.Name,
		Surname: dto.Surname,
		Password: dto.Password,
		Email: dto.Email,
		Street: dto.Street, 	 
		StreetNumber: dto.StreetNumber,
		City: dto.City,
		ZipCode: dto.ZipCode,
		Country: dto.Country 
	}
}
