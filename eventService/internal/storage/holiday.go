package storage

import (
	"eventService/internal/db"
	"eventService/internal/models"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetPrazdnikDetails(idZayavki uint8) models.Holiday {
	var prazdnik models.Prazdnik
	var ploshadki []models.Ploshadka
	var statusi []models.StatusZayavki
	var vedushie []models.Vedushiy

	// Выполняем выборку с подгрузкой связанных данных
	db.DB().Preload("Zayavka").
		Preload("Zayavka.Sotrudnik").
		Preload("Zayavka.Zakazchik").
		Preload("Zayavka.Zakazchik.StatusZakazchika").
		Preload("Zayavka.VidiPrazdnikov").
		Preload("Vedushiy").
		Preload("Ploshadka").
		Preload("Ploshadka.Address").
		Preload("Zayavka.StatusZayavki").
		First(&prazdnik, "id_zayavki = ?", idZayavki)

	db.DB().Find(&ploshadki)
	db.DB().Find(&statusi)
	db.DB().Find(&vedushie)

	var uslugi []models.UslugiVZayavkah
	db.DB().Preload("Usluga").
		Preload("Zayavka").
		Find(&uslugi, "id_zayavki = ?", idZayavki)

	var holiday models.Holiday
	holiday.Prazdnik = prazdnik
	holiday.Ploshadki = ploshadki
	holiday.Statusi = statusi
	holiday.Vedushie = vedushie
	holiday.Uslugi = uslugi

	return holiday
}

func SaveHoliday(hol models.HolidayData) bool {
	// Преобразование данных
	idZayavki, err := strconv.Atoi(hol.IDZayavki) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("error converting id_zayavki to int: %v", err)
	}

	idStatusa, err := strconv.Atoi(hol.IDStatusaZayavki) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("error converting id_statusa_zayavki to int: %v", err)
	}

	idPloshadki, err := strconv.Atoi(hol.IDPloshadki) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("error converting id_ploshadki to int: %v", err)
	}

	idVedushego, err := strconv.Atoi(hol.IDVedushego) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("error converting id_vedushego to int: %v", err)
	}

	err = callUpdateZayavkaData(idZayavki, idStatusa, idPloshadki, idVedushego, hol.DopUslugi)
	if err != nil {
		log.Println("error calling update_zayavka_data function: %v", err)
		return false
	}

	log.Printf("В заявку №%d внесены изменения idStatus=%d, idVedushego=%d, idPloshadki=%d, idDopUslug=%s изменено на true",
		idZayavki, idStatusa, idVedushego, idPloshadki, strings.Join(hol.DopUslugi, ","))

	return true
}

func callUpdateZayavkaData(idZayavki, idStatusaZayavki, idPloshadki, idVedushego int, dopUslugi []string) error {
	dopUslugiStr := fmt.Sprintf("{%v}", strings.Join(dopUslugi, ","))

	sql := `SELECT update_zayavka_data(
		?, ?, ?, ?, ?::integer[])`

	if err := db.DB().Exec(sql, idZayavki, idStatusaZayavki, idPloshadki, idVedushego, dopUslugiStr).Error; err != nil {
		return err
	}

	return nil
}

func ReadyHoliday(iDZayavki string) bool {
	sql := "SELECT calculate_and_update_total_cost(?);"

	idZayavki, err := strconv.Atoi(iDZayavki) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("error converting id_zayavki to int: %v", err)
	}

	if err := db.DB().Exec(sql, idZayavki).Error; err != nil {
		log.Println("error calling update_zayavka_data function: %v", err)
		return false
	}

	log.Printf("Статус заявки №%d изменён на 'выполнена' и посчитана полная стоимость", idZayavki)
	return true
}
