package handlers

import (
	"adminService/internal/context"
	"bytes"
	"io"
	"log"
	"net/http"
	"reflect"
)

func ApiHandler(ctx *context.Context, path string) {
	methodMap, ok := apiMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !access(ctx, path) {
		if ctx.Request.Method == "GET" {
			errorPath(ctx, 403)
			return
		}
		ctx.Response.WriteHeader(403)
		return
	}

	method, ok := methodMap[path]
	if !ok {
		http.Error(ctx.Response, "Path not found", http.StatusNotFound)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}

func debugRequestBody(ctx *context.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Ошибка чтения тела запроса:", err)
	} else {
		log.Println("Тело запроса (debug):", string(body))
	}

	// Восстанавливаем тело запроса, если оно потребуется позже
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
}
