package admin

import (
	"authService/internal/db"
	"authService/internal/models"
	"fmt"
)

func GetAdmin(u map[string]string) *models.User {
	var user *models.User
	db.DB().Where("login = ? and password = ?", u["login"], u["password"]).First(&user)
	fmt.Println(user)
	return user
}
