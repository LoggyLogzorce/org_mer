package handlers

import (
	"adminService/internal/context"
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
