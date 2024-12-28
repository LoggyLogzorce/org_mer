package user

import (
	"authService/internal/db"
	"authService/internal/models"
	"log"
	"strconv"
)

func GetUser(u map[string]string) *models.User {
	var user *models.User
	db.DB().Preload("Rola").Where("login = ? and password = ?", u["login"], u["password"]).First(&user)

	return user
}

func GetUserByLogin(u map[string]string) *models.User {
	var user *models.User
	db.DB().Preload("Rola").Where("login = ?", u["login"]).First(&user)

	return user
}

func RegisterUser(u map[string]string) bool {
	var user models.User
	var role models.RolaPolzovatelya

	err := db.DB().First(&role, "naimenovanie_roli = ?", "zakazchik").Error
	if err != nil {
		log.Println(err)
		return false
	}

	user.Login = u["login"]
	user.Password = u["password"]
	user.IDRoli = role.IDRoli

	err = db.DB().Create(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}

	err = CreateUser(u, user)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func CreateUser(zak map[string]string, user models.User) error {
	var zakazchik models.Zakazchik
	status, err := strconv.ParseUint(zak["status"], 10, 8)
	if err != nil {
		log.Println(err)
		return err
	}
	statusUint8 := uint8(status)

	zakazchik.Familiya = zak["last_name"]
	zakazchik.Imya = zak["first_name"]
	zakazchik.Otchestvo = zak["middle_name"]
	zakazchik.IdPolzovatelya = user.IDPolzovatelya
	zakazchik.Email = zak["login"]
	zakazchik.Telephone = zak["telephone"]
	zakazchik.IdStatusaZakazchika = statusUint8

	err = db.DB().Table("zakazchiki").Create(&zakazchik).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
