package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/storage"
	"log"
)

func (h *Handler) GetEvents(ctx *context.Context) {
	events := storage.GetAllEvents()

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&events)
	if err != nil {
		log.Println(err)
	}
	return
}
