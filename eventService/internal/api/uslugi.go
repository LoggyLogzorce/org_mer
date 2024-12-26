package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/storage"
	"log"
	"net/http"
)

func (h *Handler) GetUslugi(ctx *context.Context) {
	dopUslugi := storage.GetDopUslugi()

	if err := json.NewEncoder(ctx.Response).Encode(dopUslugi); err != nil {
		log.Println(err)
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}
}
