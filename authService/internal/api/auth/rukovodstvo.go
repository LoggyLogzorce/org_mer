package auth

import (
	"authService/internal/api/token"
	"authService/internal/context"
	"authService/internal/storage/user"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) AuthRukovodstvo(ctx *context.Context) {
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
	if u.IDPolzovatelya != 0 && u.Rola.NaimenovanieRoli == "rukovodstvo" {
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

	ctx.Response.WriteHeader(401)
	return
}

func (h *Handler) AuthRukByToken(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	res = &response{
		Ok: false,
	}

	if token.IsTokenValid(data["token"], "rukovodstvo") {
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
