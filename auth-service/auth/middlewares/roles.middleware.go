package middlewares

import (
	. "auth_service/auth/model"
	. "auth_service/shared"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func RolesMiddleware(roles []UserRole) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, r *http.Request) {
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
			h.ServeHTTP(rw, r)
		}

		return http.HandlerFunc(fn)
	}
}

func containsRole(roles []UserRole, role int) bool {
	for _, r := range roles {
		if int(r) == role {
			return true
		}
	}
	return false
}
