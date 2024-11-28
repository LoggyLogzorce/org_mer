package user

import (
	"authService/internal/db"
	"authService/internal/models"
)

func GetUser(u map[string]string) *models.User {
	var user *models.User
	db.DB().
		Select("polzovateli.id_polzovatelya, polzovateli.login, polzovateli.password, roli.naimenovanie_roli").
		Where("login = ? and password = ?", u["login"], u["password"]).
		Joins("INNER JOIN roli_polzovateley as roli ON polzovateli.id_roli = roli.id_roli").First(&user)

	return user
}
