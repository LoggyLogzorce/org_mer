package handler

import (
	"adminService/internal/api"
	"adminService/internal/configs"
	"adminService/internal/context"
	"adminService/internal/user"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var types map[string]bool
var cfg *configs.Config

var uHdl *user.Handler
var urlMap map[string]map[string]reflect.Value

var apiHdl *api.Handler
var apiMap map[string]map[string]reflect.Value

func init() {
	types = make(map[string]bool)
	types[".css"] = true
	types[".js"] = true
	types[".ico"] = true
	types[".jpg"] = true
	types[".png"] = true

	cfg = configs.Get()
	urlMap = make(map[string]map[string]reflect.Value)
	urlMap["POST"] = make(map[string]reflect.Value)
	urlMap["PUT"] = make(map[string]reflect.Value)
	urlMap["DELETE"] = make(map[string]reflect.Value)
	urlMap["GET"] = make(map[string]reflect.Value)

	apiMap = make(map[string]map[string]reflect.Value)
	apiMap["POST"] = make(map[string]reflect.Value)
	apiMap["PUT"] = make(map[string]reflect.Value)
	apiMap["DELETE"] = make(map[string]reflect.Value)
	apiMap["GET"] = make(map[string]reflect.Value)

	mapsHdl := cfg.Handlers
	mapsApi := cfg.Api

	uHdl = &user.Handler{}
	apiHdl = &api.Handler{}
	structHdl := reflect.TypeOf(uHdl)
	structApi := reflect.TypeOf(apiHdl)

	for methodNum := 0; methodNum < structHdl.NumMethod(); methodNum++ {
		method := structHdl.Method(methodNum)
		val, ok := mapsHdl[method.Name]
		if !ok {
			continue
		}

		urlMap[val.Method][val.Url] = reflect.ValueOf(uHdl).MethodByName(method.Name)
	}
	log.Println("urlMap has been read")

	for methodNum := 0; methodNum < structApi.NumMethod(); methodNum++ {
		method := structApi.Method(methodNum)
		val, ok := mapsApi[method.Name]
		if !ok {
			continue
		}

		apiMap[val.Method][val.Url] = reflect.ValueOf(apiHdl).MethodByName(method.Name)
	}
	log.Println("apiMap has been read")
}

func static(path string) bool {
	splitPath := strings.Split(path, "/")
	fileName := splitPath[len(splitPath)-1]
	splitName := strings.Split(fileName, ".")
	fileExt := "." + splitName[len(splitName)-1]
	if types[fileExt] {
		return true
	}
	return false
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Response: w,
		Request:  r,
	}

	path := r.URL.Path[1:]

	log.Println("Page:", r.URL.Path)
	if ok := static(path); ok {
		http.ServeFile(ctx.Response, ctx.Request, "./internal/static/"+path)
		return
	}

	pathArr := strings.Split(path, "/")
	if pathArr[0] != "api" {
		UrlHandler(ctx, path)
		return
	}

	ApiHandler(ctx, path)
	return
}

func UrlHandler(ctx *context.Context, path string) {
	methodMap, ok := urlMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !access(ctx, path) {
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

func ApiHandler(ctx *context.Context, path string) {
	methodMap, ok := apiMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !access(ctx, path) {
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

func access(ctx *context.Context, path string) bool {
	for _, value := range cfg.AccessExceptions.List {
		if value == path {
			return api.AuthByToken(ctx)
		}
	}
	return true
}
