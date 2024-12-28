package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"userService/internal/context"
	"userService/internal/models"
)

type response struct {
	Uid   uint8  `json:"uid"`
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

func (h *Handler) SendApplications(ctx *context.Context) {
	var data models.SendApplication
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		log.Println(err)
		http.Error(ctx.Response, "Bad request", http.StatusBadRequest)
		return
	}

	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println("Ошибка получения токена из куки: ", err)
		return
	}

	tokenMap := map[string]string{"token": token.Value}
	dataJson, _ := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/get/customer/id", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer res.Body.Close()

	resp := &response{}
	if err := json.NewDecoder(res.Body).Decode(resp); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	data.IdPolzovatelya = resp.Uid

	dataJson, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	responce, err := http.Post("http://localhost:8082/save/application", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer responce.Body.Close()

	if err = json.NewDecoder(responce.Body).Decode(resp); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	ctx.Response.WriteHeader(201)
}

func (h *Handler) GetApplications(ctx *context.Context) {
	var zayavki models.CustomerApplication

	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println("Ошибка получения токена из куки: ", err)
		return
	}

	tokenMap := map[string]string{"token": token.Value}
	dataJson, _ := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/get/customer/id", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer res.Body.Close()

	resp := &response{}
	if err := json.NewDecoder(res.Body).Decode(resp); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	url := fmt.Sprintf("http://localhost:8082/get/customer/applications?uid=%d", int(resp.Uid))

	responce, err := http.Get(url)
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer responce.Body.Close()

	if err = json.NewDecoder(responce.Body).Decode(&zayavki); err != nil {
		log.Println(err)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.Response).Encode(zayavki)
	if err != nil {
		log.Println(err)
		return
	}
}
