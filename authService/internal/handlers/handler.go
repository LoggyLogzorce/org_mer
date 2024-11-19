package handlers

import (
	"authService/internal/api/admin"
	"authService/internal/configs"
	"authService/internal/context"
	"log"
	"net/http"
	"reflect"
)

var hdl *admin.Handler
var apiMap map[string]map[string]reflect.Value

func init() {
	cfg := configs.Get()
	apiMap = make(map[string]map[string]reflect.Value)
	apiMap["POST"] = make(map[string]reflect.Value)

	maps := cfg.ApiMap

	hdl = &admin.Handler{}
	_struct := reflect.TypeOf(hdl)

	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}

		apiMap[val.Method][val.Url] = reflect.ValueOf(hdl).MethodByName(method.Name)
	}
	log.Println("apiMap has been read")
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Response: w,
		Request:  r,
	}

	methodMap, ok := apiMap[r.Method]
	if !ok {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Page:", r.URL.Path)
	path := r.URL.Path[1:]

	method, ok := methodMap[path]
	if !ok {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}
