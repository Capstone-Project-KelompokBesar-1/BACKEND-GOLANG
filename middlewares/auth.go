package middlewares

import (
	"errors"
	"ourgym/config"
	"ourgym/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	ID      uint   `json:"id"`
	IsAdmin bool   `json:"is_admin"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(user models.User, expLimit time.Duration) (string, error) {
	claims := &jwtCustomClaims{
		user.ID,
		user.IsAdmin,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * expLimit).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	cfg := config.Cfg

	tokenString, err := token.SignedString([]byte(cfg.JWT_SECRET_KEY))

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func GetJWTSecretKeyForAdmin(token *jwt.Token) (interface{}, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid or expired jwt")
	}

	if claims["is_admin"] != true {
		return nil, errors.New("invalid or expired jwt")
	}

	cfg := config.Cfg

	return []byte(cfg.JWT_SECRET_KEY), nil
}
