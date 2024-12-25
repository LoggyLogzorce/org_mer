package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rukService/internal/context"
	"rukService/internal/models"
)

type response struct {
	Uid uint8  `json:"uid"`
	App string `json:"app"`
	Ok  bool   `json:"ok"`
}

func (_ *Handler) GetApplications(ctx *context.Context) {
	var applications []models.Application

	r, err := http.Get("http://localhost:8082/get/applications?search=all")
	if err != nil {
		log.Println("Bad request", err)
		return
	}
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&applications); err != nil {
		log.Println(err)
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.Response).Encode(applications)
	if err != nil {
		log.Println(err)
		return
	}
}

func (_ *Handler) AcceptApplication(ctx *context.Context) {
	app := ctx.Request.URL.Query().Get("app")

	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println("Ошибка получения токена из куки: ", err)
		return
	}

	tokenMap := map[string]string{"token": token.Value}
	dataJson, _ := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/get/admin/id", "application/json", bytes.NewReader(dataJson))
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

	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8082/update/app/admin?id=%d&app=%s", int(resp.Uid), app)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	rs, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer rs.Body.Close()

	if err = json.NewDecoder(rs.Body).Decode(&resp); err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return
	}

	if !resp.Ok {
		ctx.Response.WriteHeader(400)
		return
	}

	ctx.Response.WriteHeader(200)
	return
}

func (_ *Handler) CancelApplication(ctx *context.Context) {
	app := ctx.Request.URL.Query().Get("app")

	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println("Ошибка получения токена из куки: ", err)
		return
	}

	tokenMap := map[string]string{"token": token.Value}
	dataJson, _ := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/get/admin/id", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println("Ошибка выполнения POST запроса: ", err)
		ctx.Response.WriteHeader(400)
		return
	}
	defer res.Body.Close()

	resp := &response{}
	if err = json.NewDecoder(res.Body).Decode(resp); err != nil {
		log.Println("Ошибка декодирования json1: ", err)
		ctx.Response.WriteHeader(400)
		return
	}

	url := fmt.Sprintf("http://localhost:8082/cancel/app/admin?id=%d&app=%s", int(resp.Uid), app)
	req, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println("Ошибка выполнения запроса: ", err)
		ctx.Response.WriteHeader(400)
		return
	}

	if err = json.NewDecoder(req.Body).Decode(&resp); err != nil {
		log.Println("Ошибка декодирования json: ", err)
		ctx.Response.WriteHeader(400)
		return
	}
	fmt.Println(resp.Ok)
	if !resp.Ok {
		ctx.Response.WriteHeader(400)
		return
	}

	ctx.Response.WriteHeader(200)
	return
}
