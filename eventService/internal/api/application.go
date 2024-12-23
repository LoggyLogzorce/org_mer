package api

import (
	"encoding/json"
	"eventService/internal/context"
	"eventService/internal/models"
	"eventService/internal/storage"
	"log"
	"strconv"
)

type response struct {
	Ok bool `json:"ok"`
}

func (h *Handler) GetApplications(ctx *context.Context) {
	var applications []models.Application
	var pid uint8
	search := ctx.Request.URL.Query().Get("search")
	uid := ctx.Request.URL.Query().Get("id")

	if uid != "" {
		uidUint, err := strconv.ParseUint(uid, 10, 8)
		if err != nil {
			log.Println(err)
			return
		}
		pid = uint8(uidUint)
	}

	if search == "" {
		applications = storage.GetAllApplications()
	} else if search == "my_tasks" && uid != "0" {
		sid := storage.GetSotrudnik(pid)
		applications = storage.GetMyApplications(sid)
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&applications)
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

func (h *Handler) CancelApplication(ctx *context.Context) {
	uid := ctx.Request.URL.Query().Get("id")
	app := ctx.Request.URL.Query().Get("app")

	rs := storage.CancelApplication(uid, app)

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
