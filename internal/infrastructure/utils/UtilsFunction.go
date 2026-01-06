package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string, sessionId string) (string, error) {
	expirationTime := time.Now().Add(6 * time.Hour)

	claims := &MyCustomClaims{
		UserId:    userId,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envSecret := os.Getenv("SECRET_JWT")

	tokenString, err := token.SignedString([]byte(envSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	envSecret := os.Getenv("SECRET_JWT")

	// Parsear el token usando tus claims personalizados
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validar que el método de firma sea el que esperas
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(envSecret), nil
	})

	if err != nil {
		return "", err
	}

	// Extraer los datos si el token es válido
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", jwt.ErrSignatureInvalid
}
