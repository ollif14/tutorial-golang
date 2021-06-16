package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	GenerateToken() string
}

type AuthCustomClaims struct {
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: GetSecretKey(),
		issure:    "golang-echo",
	}
}

func GetSecretKey() string {
	return "golang-echo-secret-key"
}

func (service *jwtServices) GenerateToken() string {
	claims := &AuthCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
