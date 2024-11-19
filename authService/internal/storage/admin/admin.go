package admin

import (
	"authService/internal/db"
	"authService/internal/models"
)

func GetAdmin(u map[string]string) uint8 {
	var admin *models.Admin
	db.DB().Where("login = ? and password = ?", u["login"], u["password"]).First(&admin)

	return admin.Uid
}
