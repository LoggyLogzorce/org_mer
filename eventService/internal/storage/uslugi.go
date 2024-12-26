package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
)

func GetDopUslugi() []models.Usluga {
	var dopUslugi []models.Usluga

	db.DB().Find(&dopUslugi)

	return dopUslugi
}
