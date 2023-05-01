package auth

import (
	. "auth_service/auth/model"
	. "auth_service/shared"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository IUserRepository
}

func (authService *AuthService) Register(user User) (User, *Error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, RegistrationFailed()
	}
	user.Password = hashedPassword

	fmt.Println("Hit")

	id, err := authService.UserRepository.Create(user)
	if err != nil {
		return User{}, RegistrationFailed()
	}

	created, err := authService.UserRepository.FindById(id)
	if err != nil {
		return User{}, UserDoesntExist()
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

func (authService *AuthService) GetCurrentUser(userId string) (User, *Error) {
	user, err := authService.UserRepository.FindById(StringToObjectId(userId))
	if err != nil {
		return user, InvalidCredentials()
	}
	return user, nil
}

func (authService *AuthService) DecryptToken(bearerToken string) (string, *Error) {
	token, err := validateToken(bearerToken)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", TokenValidationFailed()
	}
	id := claims["id"].(string)
	return id, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user User) (string, *Error) {
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

func (authService *AuthService) Delete(id primitive.ObjectID) *Error {
	err := authService.UserRepository.Delete(id)
	if err != nil {
		return DeleteProfileError()
	}
	return nil
}

func validateToken(bearerToken string) (*jwt.Token, *Error) {
	bearer := strings.Split(bearerToken, " ")
	token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		var secretKey = []byte(os.Getenv("JWT_SECRET"))
		return secretKey, nil
	})
	if err != nil {
		return nil, TokenValidationFailed()
	}
	return token, nil
}
