package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
)

func GetSotrudnik(idPol uint8) uint8 {
	var sotrudnik models.Sotrudnik

	db.DB().Where("id_polzovatelya = ?", idPol).First(&sotrudnik)

	return sotrudnik.IdSotrudnika
}
