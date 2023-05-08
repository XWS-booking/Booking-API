package middlewares

import (
	. "gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func UserMiddleware(next runtime.HandlerFunc) runtime.HandlerFunc {
	return runtime.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		var token *jwt.Token = context.Get(r, "Token").(*jwt.Token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			Unauthorized(rw)
			return
		}
		id := claims["id"].(string)
		context.Set(r, "id", id)
		next(rw, r, pathParams)
	})
}
