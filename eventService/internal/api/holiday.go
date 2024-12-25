package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/models"
	"eventService/internal/storage"
	"log"
	"strconv"
)

type responce struct {
	Ok bool `json:"ok"`
}

func (h *Handler) GetPrazdnik(ctx *context.Context) {
	var hol models.Holiday
	var idApp uint8
	app := ctx.Request.URL.Query().Get("app")
	if app != "" {
		uidUint, err := strconv.ParseUint(app, 10, 8)
		if err != nil {
			log.Println(err)
			return
		}
		idApp = uint8(uidUint)
	}

	hol = storage.GetPrazdnikDetails(idApp)

	err := json.NewEncoder(ctx.Response).Encode(&hol)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) SaveHoliday(ctx *context.Context) {
	var hol models.HolidayData
	if err := json.NewDecoder(ctx.Request.Body).Decode(&hol); err != nil {
		log.Println(err)
		return
	}

	ok := storage.SaveHoliday(hol)

	if ok && hol.IDStatusaZayavki == "3" {
		ok = storage.ReadyHoliday(hol.IDZayavki)
	}

	resp := responce{Ok: ok}

	if err := json.NewEncoder(ctx.Response).Encode(resp); err != nil {
		log.Println(err)
	}
}
