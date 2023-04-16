package auth

import (
	. "auth_service/auth/model"
	. "auth_service/shared"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository IUserRepository
}

func (authService *AuthService) Register(user UserModel) (UserModel, *Error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return UserModel{}, RegistrationFailed()
	}
	user.Password = hashedPassword

	id, err := authService.UserRepository.Create(user)
	if err != nil {
		return UserModel{}, RegistrationFailed()
	}

	created, err := authService.UserRepository.FindById(id)
	if err != nil {
		return UserModel{}, UserDoesntExist()
	}
	return created, nil
}

func (authService *AuthService) SignIn(email string, password string) (string, *Error) {
	user, err := authService.UserRepository.FindByEmail(email)
	if err != nil {
		return "", InvalidCredentials()
	}
	isPasswordValid := CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		return "", InvalidCredentials()
	}
	return generateToken(user)
}

func (authService *AuthService) GetCurrentUser(userId string) (UserModel, *Error) {
	user, err := authService.UserRepository.FindById(StringToObjectId(userId))
	if err != nil {
		return user, InvalidCredentials()
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user UserModel) (string, *Error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()
	claims["role"] = user.Role
	claims["id"] = user.Id
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", TokenGenerationFailed()
	}

	return tokenString, nil
}
