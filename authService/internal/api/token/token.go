package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var secretKey = []byte("rukovodstvo")

func CreateToken(uid uint8, role string) string {
	claims := jwt.MapClaims{
		"uid":  uid,
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

func IsTokenValid(tokenString string, role string) bool {
	// Парсинг токена
	token, err := ParseToken(tokenString)
	if err != nil {
		log.Println("Ошибка парсинга токена:", err)
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	if !token.Valid {
		return false
	}

	userRole, ok := claims["role"]
	if !ok || userRole != role {
		return false
	}

	return true
}

func GetUidByToken(tokenString string) (uint8, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		log.Println("Ошибка парсинга токена: ", err)
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Ошибка чтения claims")
		return 0, err
	}

	userId, ok := claims["uid"]
	if !ok {
		log.Println("Ошибка получения uid")
		return 0, err
	}

	return uint8(userId.(float64)), nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
