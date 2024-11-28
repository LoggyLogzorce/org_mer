package token

import (
	"fmt"
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

func IsTokenValid(tokenString string) bool {
	// Парсинг токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	if !token.Valid {
		return false
	}

	// Проверка наличия поля "role" и его типа
	userRole, ok := claims["role"]
	if !ok || userRole != "sotrudnik" {
		return false
	}

	return true
}
