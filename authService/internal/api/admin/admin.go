package admin

import (
	"authService/internal/api/token"
	"authService/internal/context"
	"authService/internal/storage/user"
	"encoding/json"
	"log"
	"net/http"
)

var res *response

type response struct {
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

func (h *Handler) AuthAdmin(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	res = &response{
		Ok:    false,
		Token: "",
	}

	u := user.GetUser(data)
	if u.IdPolzovatelya != 0 && u.NaimenovanieRoli == "sotrudnik" {
		tokenString := token.CreateToken(u.IdPolzovatelya, u.NaimenovanieRoli)

		res = &response{
			Ok:    true,
			Token: tokenString,
		}
		ctx.Response.Header().Set("Authorization", tokenString)
		ctx.Response.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(ctx.Response).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}

	ctx.Response.WriteHeader(401)
	return
}

func (h *Handler) AuthAdminByToken(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	res = &response{
		Ok: false,
	}

	if token.IsTokenValid(data["token"]) {
		res = &response{
			Ok: true,
		}

		ctx.Response.Header().Set("Authorization", "true")
		ctx.Response.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(ctx.Response).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}

	ctx.Response.WriteHeader(401)
	return
}
