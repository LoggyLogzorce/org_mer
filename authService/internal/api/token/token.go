package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("secret_key")

func CreateToken(uuid uint8, role string) string {
	claims := jwt.MapClaims{
		"uid":  uuid,
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}

	// Создать токен с указанными данными и алгоритмом подписи
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписать токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return tokenString
}
