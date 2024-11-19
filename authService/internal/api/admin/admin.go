package admin

import (
	"authService/internal/api/token"
	"authService/internal/context"
	"authService/internal/storage/admin"
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

	uuid := admin.GetAdmin(data)
	if uuid != 0 {
		tokenString := token.CreateToken(uuid, "admin")

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
