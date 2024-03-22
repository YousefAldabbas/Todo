package utils

import (
	"github.com/rs/zerolog/log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func GenerateToken(sub string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error().Err(err).Msg("Error signing token")
		return "", err
	}

	return tokenString, nil
}

func GeneteateRefreshToken(sub string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error().Err(err).Msg("Error signing token")
		return "", err
	}

	return tokenString, nil
}


func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Error parsing token")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error().Msg("Invalid token")
		return nil, err
	}

	return claims, nil
}