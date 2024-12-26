package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/storage"
	"log"
)

func (h *Handler) GetVidiPrazdnikov(ctx *context.Context) {
	vidiPrazdnikov := storage.GetVidiPrazdnikov()

	if err := json.NewEncoder(ctx.Response).Encode(&vidiPrazdnikov); err != nil {
		log.Println("Ошибка енкода:", err)
		return
	}
}
