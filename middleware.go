package main

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(jwtClaims jwt.StandardClaims) (string, error) {
	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return unsignedJWT.SignedString([]byte(JWT_SECRET))
}

func GenerateBackdoorJWT() (string, error) {
	return GenerateJWT(jwt.StandardClaims{
		Audience: "backdoor",
		IssuedAt: time.Now().Unix(),
	})
}

func GenerateUserJWT(userId string) (string, error) {
	return GenerateJWT(jwt.StandardClaims{
		Audience: "user",
		IssuedAt: time.Now().Unix(),
		Id:       userId,
	})
}

func VerifyJWT(jwtString string) (*jwt.Token, error) {
	jwt, err := jwt.ParseWithClaims(jwtString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if jwt == nil {
		return nil, errors.New("unable to parse JWT")
	}
	if jwt.Valid {
		return jwt, err
	} else {
		return jwt, errors.New("the JWT is not valid")
	}
}

func VerifyJWTCookie(c *gin.Context) (*jwt.Token, error) {
	jwtString, err := c.Cookie("jwt")
	if err != nil {
		return nil, err
	}
	return VerifyJWT(jwtString)
}

func SetUserJWTCookie(c *gin.Context, userID string) {
	// Create a new user JWT and sign it
	jwtString, err := GenerateUserJWT(userID)
	if err != nil {
		HandleError(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	httpCookie := &http.Cookie{
		Name:     "jwt",
		Value:    url.QueryEscape(jwtString),
		MaxAge:   int(time.Now().Unix()) * 2,
		Domain:   COOKIE_DOMAIN,
		Secure:   COOKIE_SECURE,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
	}

	// Attach the signed JWT as a secure, httpOnly cookie
	http.SetCookie(c.Writer, httpCookie)
}
