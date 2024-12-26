package api

import (
	"encoding/json"
	"log"
	"net/http"
	"userService/internal/context"
	"userService/internal/models"
)

func (h *Handler) GetStatusi(ctx *context.Context) {
	var statusi []models.StatusZakazchika

	res, err := http.Get("http://localhost:8082/get/statusi-zakazchikov")
	if err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request", http.StatusBadRequest)
		return
	}

	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&statusi); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Error reading body", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(ctx.Response).Encode(statusi); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Error encode json", http.StatusBadRequest)
		return
	}
}
