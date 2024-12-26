package user

import (
	"net/http"
	"time"
	"userService/internal/context"
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
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/login.html")
	return
}

func (h *Handler) HomePage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index_auth.html")
	return
}

func (h *Handler) Home(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index.html")
	return
}

func (h *Handler) Uslugi(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/uslugi.html")
	return
}

func (h *Handler) RegisterPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/register.html")
	return
}
