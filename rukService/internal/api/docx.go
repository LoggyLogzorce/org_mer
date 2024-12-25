package api

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"rukService/internal/context"
	"strings"
	"time"
)

func (h *Handler) CreateDocx(ctx *context.Context) {
	appId := ctx.Request.URL.Query().Get("app")

	application := GetPrazdnik(appId)

	// Открываем шаблон DOCX
	r, err := docx.ReadDocxFile("./internal/docx/template/template.docx")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer r.Close()

	var uslugiBuilder strings.Builder
	for _, v := range application.Uslugi {
		if v.Complete {
			uslugiBuilder.WriteString(fmt.Sprintf("%s - %.2f, ", v.Usluga.Naimenovanie, v.Usluga.Stoimost))
		}
	}
	uslugi := uslugiBuilder.String()
	if len(uslugi) != 0 {
		uslugi = uslugi[:len(uslugi)-2]
	} else {
		uslugi = "Дополнительные услуги не выбраны."
	}

	summaApp := fmt.Sprintf("%v", application.Prazdnik.Zayavka.VidiPrazdnikov.Summa)
	summaPloshadka := fmt.Sprintf("%v", application.Prazdnik.Ploshadka.StoimostVChas)
	stavkaVedushiy := fmt.Sprintf("%v", application.Prazdnik.Vedushiy.StavkaVChas)
	commisiyaSotr := application.Prazdnik.PolnayaStoimost * 2.0 / 102.0
	commisiya := fmt.Sprintf("%v", commisiyaSotr)
	fullPrice := fmt.Sprintf("%v", application.Prazdnik.PolnayaStoimost)
	idApp := fmt.Sprintf("%v", application.Prazdnik.IdZayavki)

	// Получаем содержимое документа
	doc := r.Editable()
	doc.Replace("{{idZayavki}}", idApp, -1)

	doc.Replace("{{FamilyaZ}}", application.Prazdnik.Zayavka.Zakazchik.Familiya, -1)
	doc.Replace("{{ImyaZ}}", application.Prazdnik.Zayavka.Zakazchik.Imya, -1)
	doc.Replace("{{OtchestvoZ}}", application.Prazdnik.Zayavka.Zakazchik.Otchestvo, -1)
	doc.Replace("{{StatusZ}}", application.Prazdnik.Zayavka.Zakazchik.StatusZakazchika.NaimenovanieStatusa, -1)
	doc.Replace("{{PhoneNumberZ}}", application.Prazdnik.Zayavka.Zakazchik.Telephone, -1)
	doc.Replace("{{EmailZ}}", application.Prazdnik.Zayavka.Zakazchik.Email, -1)

	doc.Replace("{{typeApp}}", application.Prazdnik.Zayavka.VidiPrazdnikov.NaimenovanieVida, -1)
	doc.Replace("{{Price}}", summaApp, -1)
	doc.Replace("{{Date}}", strings.Split(application.Prazdnik.Zayavka.DataProvedeniya, "T")[0], -1)
	doc.Replace("{{nachalo}}", application.Prazdnik.Zayavka.NachaloProvedeniya, -1)
	doc.Replace("{{konec}}", application.Prazdnik.Zayavka.KonecProvedeniya, -1)

	doc.Replace("{{uslugi}}", uslugi, -1)

	doc.Replace("{{NamePlace}}", application.Prazdnik.Ploshadka.Nazvanie, -1)
	doc.Replace("{{Gorod}}", application.Prazdnik.Ploshadka.Address.Gorod, -1)
	doc.Replace("{{Ulica}}", application.Prazdnik.Ploshadka.Address.Ulica, -1)
	doc.Replace("{{Dom}}", application.Prazdnik.Ploshadka.Address.Dom, -1)
	doc.Replace("{{Korpus}}", application.Prazdnik.Ploshadka.Address.Korpus, -1)
	doc.Replace("{{stoimost}}", summaPloshadka, -1)

	doc.Replace("{{FamilyaV}}", application.Prazdnik.Vedushiy.Familiya, -1)
	doc.Replace("{{ImyaV}}", application.Prazdnik.Vedushiy.Imya, -1)
	doc.Replace("{{OtchestvoV}}", application.Prazdnik.Vedushiy.Otchestvo, -1)
	doc.Replace("{{PhoneNumberV}}", application.Prazdnik.Vedushiy.Telephone, -1)
	doc.Replace("{{stavkaV}}", stavkaVedushiy, -1)

	doc.Replace("{{FamilyaS}}", application.Prazdnik.Zayavka.Sotrudnik.Familiya, -1)
	doc.Replace("{{ImyaS}}", application.Prazdnik.Zayavka.Sotrudnik.Imya, -1)
	doc.Replace("{{OtchestvoS}}", application.Prazdnik.Zayavka.Sotrudnik.Otchestvo, -1)
	doc.Replace("{{PhoneNumberS}}", application.Prazdnik.Zayavka.Sotrudnik.Telephone, -1)
	doc.Replace("{{EmailS}}", application.Prazdnik.Zayavka.Sotrudnik.Email, -1)

	doc.Replace("{{KomSot}}", commisiya, -1)

	doc.Replace("{{FullPrice}}", fullPrice, -1)

	// Сохраняем изменения
	timeDate := time.Now()
	toDayDate := timeDate.Format("2006-01-02")
	pathToFile := fmt.Sprintf("./internal/docx/otchet_zayavka№%s_%s.docx", idApp, toDayDate)
	err = doc.WriteToFile(pathToFile)
	if err != nil {
		log.Fatalf("Ошибка сохранения файла: %v", err)
	}

	log.Println("Файл успешно изменён:", pathToFile)
}

func (h *Handler) DownLoadReport(ctx *context.Context) {
	app := ctx.Request.URL.Query().Get("app")

	timeDate := time.Now()
	toDayDate := timeDate.Format("2006-01-02")
	pathToFile := fmt.Sprintf("./internal/docx/otchet_zayavka№%s_%s.docx", app, toDayDate)

	// Открываем файл
	file, err := os.Open(pathToFile)
	if err != nil {
		http.Error(ctx.Response, "Файл не найден", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Получаем базовое имя файла
	fileName := filepath.Base(pathToFile)

	// Устанавливаем заголовки для загрузки файла
	ctx.Response.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	ctx.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

	// Отправляем файл пользователю
	http.ServeFile(ctx.Response, ctx.Request, pathToFile)
}
