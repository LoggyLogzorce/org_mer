package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"userService/internal/context"
)

type request struct {
	Ok bool `json:"ok"`
}

func (h *Handler) Auth(ctx *context.Context) {
	var data map[string]string
	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}

	dataJson, err := json.Marshal(data)

	res, err := http.Post("http://localhost:8081/auth/user", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(ctx.Request.Body).Decode(res)
	token := res.Header.Get("Authorization")

	if token != "" {
		cookie := &http.Cookie{
			Name:  "token",
			Value: token,
			Path:  "/",
		}
		http.SetCookie(ctx.Response, cookie)

		response := struct {
			Ok bool `json:"ok"`
		}{
			Ok: true,
		}
		ctx.Response.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(ctx.Response).Encode(response)
		return
	}

	ctx.Response.WriteHeader(401)
	return
}

func AuthByToken(ctx *context.Context) bool {
	// Чтение тела запроса и сохранение в переменной
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Ошибка чтения тела запроса:", err)
		return false
	}

	// Восстановление тела запроса
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println(err)
		return false
	}

	tokenMap := map[string]string{"token": token.Value}

	dataJson, err := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/auth/user/token", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return false
	}

	err = json.NewDecoder(res.Body).Decode(res)
	val := res.Header.Get("Authorization")

	if val != "" {
		return true
	}

	return false
}

func (h *Handler) Register(ctx *context.Context) {
	var data map[string]string
	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}

	req := request{Ok: false}

	dataJson, err := json.Marshal(data)

	res, err := http.Post("http://localhost:8081/register/user", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println("Ошибка запроса", err)
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&req); err != nil {
		log.Println("Ошибка декода", err)
		http.Error(ctx.Response, "Bad request: error decode json", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(ctx.Response).Encode(req); err != nil {
		log.Println("Ошибка енкода", err)
		http.Error(ctx.Response, "Bad request: error encode json", http.StatusBadRequest)
		return
	}
}
