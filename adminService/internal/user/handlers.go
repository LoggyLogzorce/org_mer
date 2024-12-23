package user

import (
	"adminService/internal/context"
	"net/http"
	"time"
)

type Handler struct {
}

func SetHeaders(ctx *context.Context) {
	ctx.Response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response.Header().Set("Expires", time.Now().Format(http.TimeFormat))
	ctx.Response.Header().Set("Pragma", "no-cache")
}

func (h *Handler) LoginPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index.html")
	return
}

func (h *Handler) HomePage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/homepage.html")
	return
}

func (h *Handler) TasksPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/tasks.html")
	return
}
