package auth

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
	Uid   uint8  `json:"uid"`
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

func (h *Handler) AuthAdmin(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	u := user.GetUser(data)
	if u.IDPolzovatelya != 0 && u.Rola.NaimenovanieRoli == "sotrudnik" {
		tokenString := token.CreateToken(u.IDPolzovatelya, u.Rola.NaimenovanieRoli)

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

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&res)
	if err != nil {
		log.Println(err)
	}
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

	if token.IsTokenValid(data["token"], "sotrudnik") {
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

func (h *Handler) GetAdminIdByToken(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	res = &response{
		Uid: 0,
		Ok:  false,
	}

	uid, err := token.GetUidByToken(data["token"])
	ctx.Response.Header().Set("Content-Type", "application/json")
	if err != nil {
		err := json.NewEncoder(ctx.Response).Encode(&res)
		if err != nil {
			log.Println(err)
		}

	}

	res = &response{
		Uid: uid,
		Ok:  true,
	}
	err = json.NewEncoder(ctx.Response).Encode(&res)
	return
}
