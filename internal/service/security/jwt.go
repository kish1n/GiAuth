package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/data"
	"time"
)

func GenerateJWT(user *data.User, cfg config.Config) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * cfg.JWT().ExpirationTime).Unix(),
		Subject:   user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWT().SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
