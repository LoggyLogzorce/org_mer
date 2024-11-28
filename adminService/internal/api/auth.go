package api

import (
	"adminService/internal/context"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) Auth(ctx *context.Context) {
	var data map[string]string
	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	dataJson, err := json.Marshal(data)

	res, err := http.Post("http://localhost:8081/auth/admin", "application/json", bytes.NewReader(dataJson))
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
		err = json.NewEncoder(ctx.Response).Encode(response)
		return
	}

	ctx.Response.WriteHeader(401)
	return
}

func AuthByToken(ctx *context.Context) bool {
	token, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Println(err)
		return false
	}

	tokenMap := map[string]string{"token": token.Value}

	dataJson, err := json.Marshal(tokenMap)

	res, err := http.Post("http://localhost:8081/auth/admin/token", "application/json", bytes.NewReader(dataJson))
	if err != nil {
		log.Println(err)
		ctx.Response.WriteHeader(400)
		return false
	}

	err = json.NewDecoder(ctx.Request.Body).Decode(res)
	val := res.Header.Get("Authorization")

	fmt.Println(val)

	if val != "" {
		return true
	}

	return false
}
