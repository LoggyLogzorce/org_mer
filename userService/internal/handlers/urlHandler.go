package handlers

import (
	"log"
	"net/http"
	"reflect"
	"userService/internal/context"
)

func UrlHandler(ctx *context.Context, path string) {
	methodMap, ok := urlMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !access(ctx, path) {
		if path != "" {
			errorPath(ctx, 403)
			return
		}
		path = "homepage"
	}

	method, ok := methodMap[path]
	if !ok {
		errorPath(ctx, 404)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}
