package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
	"log"
	"strconv"
)

func GetAllApplications() []models.Application {
	var applications []models.Application
	var IdStatusaZayavki uint8

	var status []models.StatusZayavki
	db.DB().Find(&status)

	for _, v := range status {
		if v.NaimenovanieStatusa == "Не принята" {
			IdStatusaZayavki = v.IdStatusa
			break
		}
	}

	db.DB().
		Preload("Zakazchik").      // Загрузить данные из таблицы zakazchiki
		Preload("VidiPrazdnikov"). // Загрузить данные из таблицы vidi_prazdnikov
		Preload("Zakazchik.StatusZakazchika").
		Where("id_sotrudnika IS NULL and id_statusa_zayavki = ?", IdStatusaZayavki).
		Find(&applications)

	for i := range applications {
		applications[i].DataProvedeniya = applications[i].DataProvedeniya[0:10]
		applications[i].NachaloProvedeniya = applications[i].NachaloProvedeniya[0:5]
		applications[i].KonecProvedeniya = applications[i].KonecProvedeniya[0:5]
	}

	return applications
}

func GetMyApplications(uid uint8) []models.Application {
	var applications []models.Application
	var IdStatusaZayavki uint8

	var status []models.StatusZayavki
	db.DB().Find(&status)

	for _, v := range status {
		if v.NaimenovanieStatusa == "В работе" {
			IdStatusaZayavki = v.IdStatusa
			break
		}
	}

	db.DB().
		Preload("Zakazchik").      // Загрузить данные из таблицы zakazchiki
		Preload("VidiPrazdnikov"). // Загрузить данные из таблицы vidi_prazdnikov
		Preload("Zakazchik.StatusZakazchika").
		Where("id_sotrudnika = ? and id_statusa_zayavki = ?", uid, IdStatusaZayavki).
		Find(&applications)

	for i := range applications {
		applications[i].DataProvedeniya = applications[i].DataProvedeniya[0:10]
		applications[i].NachaloProvedeniya = applications[i].NachaloProvedeniya[0:5]
		applications[i].KonecProvedeniya = applications[i].KonecProvedeniya[0:5]
	}

	return applications
}

func AcceptApplication(uid, app string) bool {
	var application models.Application

	appUint8, err := strconv.ParseUint(app, 10, 8)
	if err != nil {
		log.Println(err)
		return false
	}
	idZayavki := uint8(appUint8)

	uidUint, err := strconv.ParseUint(uid, 10, 8)
	if err != nil {
		log.Println(err)
		return false
	}
	pid := uint8(uidUint)

	idS := GetSotrudnik(pid)
	application.IdSotrudnika = &idS

	var status []models.StatusZayavki
	db.DB().Find(&status)

	for _, v := range status {
		if v.NaimenovanieStatusa == "В работе" {
			application.IdStatusaZayavki = v.IdStatusa
			break
		}
	}

	result := db.DB().Table("zayavki").Where("id_zayavki = ?", idZayavki).Updates(application)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	}

	log.Println("Заявка №", app, "взята сотрудником id =", *application.IdSotrudnika)
	return true
}

func CancelApplication(uid, app string) bool {
	var application models.Application

	appUint8, err := strconv.ParseUint(app, 10, 8)
	if err != nil {
		log.Println(err)
		return false
	}
	idZayavki := uint8(appUint8)

	uidUint, err := strconv.ParseUint(uid, 10, 8)
	if err != nil {
		log.Println(err)
		return false
	}
	pid := uint8(uidUint)

	idS := GetSotrudnik(pid)

	//TODO изменить на функцию в базе
	var status models.StatusZayavki
	db.DB().Where("naimenovanie_statusa = 'Не принята'").First(&status)

	db.DB().Where("id_zayavki = ?", idZayavki).First(&application)
	if *application.IdSotrudnika != idS {
		return false
	}

	application.IdSotrudnika = nil
	application.IdStatusaZayavki = status.IdStatusa
	db.DB().Save(&application)
	log.Println("Сотрудник id =", idS, "отказался от работы над заявкой №", application.IdZayavki)
	return true
}
