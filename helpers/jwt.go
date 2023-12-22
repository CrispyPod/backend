package helpers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

type contextKey struct {
	name string
}

var jwtUserCtxKey = &contextKey{"userName"}

const userKey = "jti"
const expKey = "exp"

func GenerateJwt(userName string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = os.Getenv("JWT_SECRET")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		ID:        userName,
	})
	return token.SignedString([]byte(jwtSecret))
}

func ParseJwt(tokenString string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = os.Getenv("JWT_SECRET")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if err == nil && ok && token.Valid {
		if err != nil || time.Unix(int64(claims[expKey].(float64)), 0).Before(time.Now()) {
			return "", errors.New("invalid jwt or jwt expired")
		}
		return claims[userKey].(string), nil
	} else {
		return "", errors.New("JWT not valid")
	}
}

func JWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			return
			// context.WithValue(ctx.Request.Context(), jwtUserCtxKey, nil)
		}

		splitted := strings.SplitN(authHeader, " ", 2)
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			return
		}
		userName, err := ParseJwt(splitted[1])
		if err != nil {
			return
		}

		rc := context.WithValue(ctx.Request.Context(), jwtUserCtxKey, userName)
		ctx.Request = ctx.Request.WithContext(rc)
		ctx.Next()
	}
}

func JWTFromContext(ctx context.Context) string {
	raw, _ := ctx.Value(jwtUserCtxKey).(string)
	return raw
}
