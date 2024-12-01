package api

import (
	"adminService/internal/context"
	"adminService/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (_ *Handler) GetEvents(ctx *context.Context) {
	var events []models.Event
	//auth := AuthByToken(ctx)
	//if !auth {
	//	ctx.Response.Header().Set("Content-Type", "application/json")
	//	err := json.NewEncoder(ctx.Response).Encode(events)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	return
	//}

	resp, err := http.Get("http://localhost:8082/get/events")
	if err != nil {
		// Обработка ошибки
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&events); err != nil {
		log.Println(err)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.Response).Encode(events)
	if err != nil {
		log.Println(err)
		return
	}
}
