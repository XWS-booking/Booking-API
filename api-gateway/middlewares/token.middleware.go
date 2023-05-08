package middlewares

import (
	"fmt"
	. "gateway/shared"
	"github.com/golang-jwt/jwt/v5"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
)

func TokenValidationMiddleware(next runtime.HandlerFunc) runtime.HandlerFunc {
	return runtime.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		if r.Header["Authorization"] == nil {
			Unauthorized(rw)
			return
		}
		bearer := strings.Split(r.Header["Authorization"][0], " ")
		token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			var secretKey = []byte(os.Getenv("JWT_SECRET"))
			return secretKey, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			Unauthorized(rw)
			return
		}
		context.Set(r, "Token", token)
		next(rw, r, pathParams)
	})
}
