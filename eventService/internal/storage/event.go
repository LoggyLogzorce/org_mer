package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
	"log"
	"strconv"
)

func GetAllEvents() []models.Event {
	var events []models.Event
	// TODO исправить на выборку из базы статусов заявок и подставление в запрос для коректной работы при изменении базы
	db.DB().
		Select("zayavki.id_zayavki, vidi_prazdnikov.naimenovanie_vida, zayavki.data_provedeniya, zakazchiki.familiya, " +
			"zakazchiki.imya, zakazchiki.otchestvo, statusi_zakazchikov.naimenovanie_statusa, " +
			"zakazchiki.telephone, zakazchiki.email, zayavki.nachalo_provedeniya, zayavki.konec_provedeniya, " +
			"zayavki.prodoljitelnost, zayavki.kolichestvo_chelovek").
		Where("zayavki.id_sotrudnika IS NULL and zayavki.status_zayavki = 1").
		Joins("inner join vidi_prazdnikov on zayavki.id_vida_prazdnika = vidi_prazdnikov.id_vida " +
			"inner join zakazchiki on zayavki.id_zakazchika = zakazchiki.id_zakazchika " +
			"inner join statusi_zakazchikov on zakazchiki.id_statusa_zakazchika = statusi_zakazchikov.id_statusa").Find(&events)

	for i := range events {
		events[i].DataProvedeniya = events[i].DataProvedeniya[0:10]
		events[i].NachaloProvedeniya = events[i].NachaloProvedeniya[0:5]
		events[i].KonecProvedeniya = events[i].KonecProvedeniya[0:5]
	}

	return events
}

func AcceptApplication(uid, app string) bool {
	var event models.Event

	appUint8, err := strconv.ParseUint(app, 10, 8)
	if err != nil {
		log.Println(err)
		return false
	}

	idZayavki := uint8(appUint8)
	event.IdSotrudnika = GetSotrudnik(uid)
	// TODO заменить на динамический выбор статуса из базы
	event.IdStatusaZayavki = 2

	result := db.DB().Table("zayavki").Where("id_zayavki = ?", idZayavki).Updates(event)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	}

	log.Println("Заявка №", app, "взята сотрудником", event.IdSotrudnika)
	return true
}
