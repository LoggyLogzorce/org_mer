package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
	"log"
)

func GetStatusiZakazchikov() []models.StatusZakazchika {
	var statusi []models.StatusZakazchika
	err := db.DB().Find(&statusi).Error
	if err != nil {
		log.Println(err)
	}
	return statusi
}
