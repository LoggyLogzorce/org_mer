package api

import (
	"adminService/internal/context"
	"adminService/internal/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetPrazdnik(ctx *context.Context) {
	var hol models.Holiday
	app := ctx.Request.URL.Query().Get("id")

	url := "http://localhost:8082/get/holiday?app=" + app
	r, err := http.Get(url)
	if err != nil {
		log.Println("Bad request", err)
		return
	}
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&hol); err != nil {
		log.Println(err)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.Response).Encode(hol)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) SaveHoliday(ctx *context.Context) {
	var updateData models.HolidayData
	if err := json.NewDecoder(ctx.Request.Body).Decode(&updateData); err != nil {
		log.Println(err)
		return
	}

	dataJson, _ := json.Marshal(updateData)

	res, err := http.Post("http://localhost:8082/save/holiday", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer res.Body.Close()

	resp := response{}

	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.Println(err)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.Response).Encode(resp)
	if err != nil {
		log.Println(err)
		return
	}
}
