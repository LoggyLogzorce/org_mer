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
		applications = storage.GetAllApplicationsInWork()
	} else if search == "my_tasks" && uid != "0" {
		sid := storage.GetSotrudnik(pid)
		applications = storage.GetMyApplications(sid)
	} else if search == "all" {
		applications = storage.GetAllApplications()
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

func (h *Handler) SaveApplication(ctx *context.Context) {
	var data models.SendApplication
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	err := storage.SaveApplication(data)
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	res := response{
		Ok: true,
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(ctx.Response).Encode(&res); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	ctx.Response.WriteHeader(201)
}

func (h *Handler) GetCustomerApplications(ctx *context.Context) {
	uid := ctx.Request.URL.Query().Get("uid")

	applications := storage.GetCustomerApplications(uid)

	if err := json.NewEncoder(ctx.Response).Encode(&applications); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
}
