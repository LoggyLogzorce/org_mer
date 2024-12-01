package handlers

import (
	"adminService/internal/context"
	"log"
	"net/http"
	"reflect"
)

func UrlHandler(ctx *context.Context, path string) {
	methodMap, ok := urlMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !access(ctx, path) {
		errorPath(ctx, 403)
		return
	}

	method, ok := methodMap[path]
	if !ok {
		errorPath(ctx, 404)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}
