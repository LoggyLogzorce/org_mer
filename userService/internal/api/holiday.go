package api

import (
	"encoding/json"
	"log"
	"net/http"
	"userService/internal/context"
	"userService/internal/models"
)

func (h *Handler) GetVidiPrazdnikov(ctx *context.Context) {
	var vidiPrazdnikov []models.VidiPrazdnikov

	res, err := http.Get("http://localhost:8082/get/vidi/prazdnikov")
	if err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request: 8082", http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&vidiPrazdnikov); err != nil {
		log.Println("Ошибка енкода:", err)
		http.Error(ctx.Response, "Bad request: error decode json", http.StatusBadRequest)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(ctx.Response).Encode(vidiPrazdnikov); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request: error encode json", http.StatusBadRequest)
		return
	}
}
