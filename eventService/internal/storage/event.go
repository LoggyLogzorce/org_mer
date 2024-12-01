package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
	"fmt"
)

func GetAllEvents() []models.Event {
	var events []models.Event
	db.DB().
		Select("zayavki.id_zayavki, vidi_prazdnikov.naimenovanie_vida, zayavki.data_provedeniya, zakazchiki.familiya, " +
			"zakazchiki.imya, zakazchiki.otchestvo, statusi_zakazchikov.naimenovanie_statusa, " +
			"zakazchiki.telephone, zakazchiki.email").
		Joins("inner join vidi_prazdnikov on zayavki.id_vida_prazdnika = vidi_prazdnikov.id_vida " +
			"inner join zakazchiki on zayavki.id_zakazchika = zakazchiki.id_zakazchika " +
			"inner join statusi_zakazchikov on zakazchiki.id_statusa_zakazchika = statusi_zakazchikov.id_statusa").Find(&events)

	for i := range events {
		events[i].DataProvedeniya = events[i].DataProvedeniya[0:10]
	}
	fmt.Println(events)
	return events
}
