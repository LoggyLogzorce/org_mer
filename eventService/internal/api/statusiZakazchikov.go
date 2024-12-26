package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/storage"
	"log"
)

func (h *Handler) GetStatusiZakazchikov(ctx *context.Context) {
	statusi := storage.GetStatusiZakazchikov()

	if err := json.NewEncoder(ctx.Response).Encode(statusi); err != nil {
		log.Println(err)
		return
	}
}
