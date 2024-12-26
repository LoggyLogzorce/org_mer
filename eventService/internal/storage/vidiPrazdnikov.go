package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
)

func GetVidiPrazdnikov() []models.VidiPrazdnikov {
	var vidiPrazdnikov []models.VidiPrazdnikov

	db.DB().Find(&vidiPrazdnikov)

	return vidiPrazdnikov
}
