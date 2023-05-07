package middlewares

import (
	. "gateway/model"
	. "gateway/shared"
	"github.com/golang-jwt/jwt/v5"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"

	"github.com/gorilla/context"
)

func RolesMiddleware(roles []UserRole, next runtime.HandlerFunc) runtime.HandlerFunc {
	return runtime.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		var token *jwt.Token = context.Get(r, "Token").(*jwt.Token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			Unauthorized(rw)
			return
		}
		role := claims["role"].(float64)
		if !containsRole(roles, int(role)) {
			Forbidden(rw)
			return
		}
		next(rw, r, pathParams)
	})
}

func containsRole(roles []UserRole, role int) bool {
	for _, r := range roles {
		if int(r) == role {
			return true
		}
	}
	return false
}
