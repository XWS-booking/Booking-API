package auth

import (
	. "auth_service/auth/model"
	sagaConfig "auth_service/auth/saga-config"
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
	UserRepository         IUserRepository
	DeleteHostOrchestrator *sagaConfig.DeleteHostOrchestrator
}

func (authService *AuthService) Register(user User) (User, *Error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, RegistrationFailed()
	}
	user.Password = hashedPassword
	user.Distinguished = false
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

func (authService *AuthService) UpdatePersonalInfo(user User) (User, *Error) {
	foundUser, err := authService.UserRepository.FindById(user.Id)
	foundUser.Name = user.Name
	foundUser.Surname = user.Surname
	foundUser.Email = user.Email
	foundUser.Street = user.Street
	foundUser.StreetNumber = user.StreetNumber
	foundUser.City = user.City
	foundUser.ZipCode = user.ZipCode
	foundUser.Country = user.Country
	foundUser.Username = user.Username
	foundUser.DeleteStatus = user.DeleteStatus
	updatedUser, err := authService.UserRepository.UpdatePersonalInfo(foundUser)
	fmt.Println(err)
	if err != nil {
		return user, PersonalInfoUpdateFailed()
	}
	return updatedUser, nil
}

func (authService *AuthService) ChangePassword(id string, oldPassword string, newPassword string) *Error {
	user, e := authService.FindById(StringToObjectId(id))
	if e != nil {
		return e
	}
	isPasswordValid := CheckPasswordHash(oldPassword, user.Password)
	if !isPasswordValid {
		return InvalidCredentials()
	}
	hashedPass, err := HashPassword(newPassword)
	if err != nil {
		return PersonalInfoUpdateFailed()
	}
	user.Password = hashedPass
	_, err = authService.UserRepository.UpdatePersonalInfo(user)
	if err != nil {
		return PersonalInfoUpdateFailed()
	}
	return nil
}

func (authService *AuthService) ChangeHostDistinguishedStatus(id primitive.ObjectID) *Error {
	user, e := authService.FindById(id)
	if e != nil {
		return e
	}
	user.Distinguished = !user.Distinguished
	_, err := authService.UserRepository.UpdatePersonalInfo(user)
	if err != nil {
		return PersonalInfoUpdateFailed()
	}
	return nil
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

func (authService *AuthService) FindById(id primitive.ObjectID) (User, *Error) {
	user, err := authService.UserRepository.FindById(id)
	if err != nil {
		return user, UserNotFoundError()
	}
	return user, nil
}

func (authService *AuthService) InitiateProfileDeletion(id primitive.ObjectID) (User, *Error) {
	user, err := authService.UserRepository.FindById(id)

	if err != nil {
		return User{}, UserNotFoundError()
	}
	user.DeleteStatus = PENDING
	authService.UserRepository.UpdatePersonalInfo(user)
	fmt.Println("orch", authService.DeleteHostOrchestrator)
	err = authService.DeleteHostOrchestrator.Start(id.Hex())
	if err != nil {
		user.DeleteStatus = ACTIVE
		authService.UserRepository.UpdatePersonalInfo(user)
		return user, UserNotFoundError()
	}
	return user, nil
}

func (authService *AuthService) GetFeaturedHosts() []string {
	hosts, err := authService.UserRepository.GetFeaturedHosts()
	if err != nil {
		return []string{}
	}
	result := make([]string, 0)
	for _, host := range hosts {
		result = append(result, host.Id.Hex())
	}
	return result

}
