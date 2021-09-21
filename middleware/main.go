package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/symmetric-project/node-backend/env"
)

type ResolverContext struct {
	JWT    *string
	Writer *gin.ResponseWriter
}

func GetResolverContext(ctx context.Context) ResolverContext {
	return ctx.Value("resolverContext").(ResolverContext)
}

func GenerateJWT(claims jwt.StandardClaims) (string, error) {
	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return unsignedJWT.SignedString([]byte(env.CONFIG.JWT_SECRET))
}

func GenerateBackdoorJWT() (string, error) {
	return GenerateJWT(jwt.StandardClaims{
		Audience: "backdoor",
		IssuedAt: time.Now().Unix(),
	})
}

func GenerateUserJWT(userName string) (string, error) {
	return GenerateJWT(jwt.StandardClaims{
		Audience: "user",
		IssuedAt: time.Now().Unix(),
		Id:       userName,
	})
}

func VerifyJWT(jwtString string) (*jwt.Token, *jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(jwtString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.CONFIG.JWT_SECRET), nil
	})
	if jwtToken == nil {
		return nil, nil, errors.New("unable to parse JWT")
	}
	if jwtToken.Valid {
		return jwtToken, jwtToken.Claims.(*jwt.StandardClaims), err
	}
	return jwtToken, jwtToken.Claims.(*jwt.StandardClaims), errors.New("the JWT is not valid")
}

func SetCookie(writer gin.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(writer, cookie)
}

func NewCookie(jwt string) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    url.QueryEscape(jwt),
		MaxAge:   int(time.Now().Unix()) * 2,
		Domain:   env.CONFIG.DOMAIN,
		Secure:   env.CONFIG.SECURE_COOKIES,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
	}
}
