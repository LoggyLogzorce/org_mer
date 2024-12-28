package auth

import (
	"authService/internal/api/token"
	"authService/internal/context"
	"authService/internal/storage/user"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) AuthUser(ctx *context.Context) {
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
	if u.IDPolzovatelya != 0 && u.Rola.NaimenovanieRoli == "zakazchik" {
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

func (h *Handler) AuthUserByToken(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	res = &response{
		Ok: false,
	}

	if token.IsTokenValid(data["token"], "zakazchik") {
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

func (h *Handler) RegisterUser(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		log.Println(err)
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
	}

	userInDb := user.GetUserByLogin(data)
	if userInDb.IDPolzovatelya != 0 {
		ctx.Response.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(ctx.Response).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		http.Error(ctx.Response, "user already exists", http.StatusConflict)
		return
	}

	ok := user.RegisterUser(data)
	if !ok {
		http.Error(ctx.Response, "Error creating user", http.StatusConflict)
		return
	}

	res = &response{
		Ok: true,
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(&res)
	if err != nil {
		log.Println(err)
	}

	ctx.Response.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) GetCustomerIdByToken(ctx *context.Context) {
	var data map[string]string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusBadRequest)
		return
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
