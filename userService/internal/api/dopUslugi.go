package api

import (
	"encoding/json"
	"log"
	"net/http"
	"userService/internal/context"
	"userService/internal/models"
)

func (h *Handler) GetUslugi(ctx *context.Context) {
	var dopUslugi []models.Usluga

	resp, err := http.Get("http://localhost:8082/get/dop-uslugi")
	if err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request: 8082", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&dopUslugi); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request: error decode json", http.StatusBadRequest)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(ctx.Response).Encode(dopUslugi); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request: error encode json", http.StatusBadRequest)
		return
	}
}
