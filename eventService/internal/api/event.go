package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/storage"
	"log"
)

type response struct {
	Ok bool `json:"ok"`
}

func (h *Handler) GetEvents(ctx *context.Context) {
	events := storage.GetAllEvents()

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&events)
	if err != nil {
		log.Println(err)
	}
	return
}

func (h *Handler) AcceptApplication(ctx *context.Context) {
	uid := ctx.Request.URL.Query().Get("id")
	app := ctx.Request.URL.Query().Get("app")

	rs := storage.AcceptApplication(uid, app)

	res := response{
		Ok: rs,
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&res)
	if err != nil {
		log.Println(err)
		return
	}
}
