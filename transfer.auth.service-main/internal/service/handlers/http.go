package handlers

import (
	"fmt"
	"github.com/transferMVP/transfer.webapp/internal/pkg/antiflood"
	err "github.com/transferMVP/transfer.webapp/internal/pkg/errorsApp"
	"github.com/transferMVP/transfer.webapp/internal/pkg/response"
	"github.com/transferMVP/transfer.webapp/internal/service/logic"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

func RoutingHttp(ctx *fasthttp.RequestCtx) {
	fmt.Println("routing")
	resp := response.InitResp()
	path := string(ctx.Path())

	defer func() {
		ctx.Response.Header.Add("Content-Type", "application/json")
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowOrigin, "*")
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodPost)
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodGet)
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodPatch)
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderContentType)
		ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderAuthorization)

		if r := recover(); r != nil {
			fmt.Printf("%s\n", debug.Stack())
			fmt.Println(fmt.Sprint(r))
			resp.SetError(err.GetErr(2, fmt.Sprint(r)))
		}
		ctx.Write(resp.FormResponse().Json())
	}()
	if err := antiflood.FastHttpCounter(ctx); err != nil {
		return
	}
	fmt.Println(path)
	switch path {
	case "/registr":
		logic.RegistrUser(resp, ctx)
	case "/getuser":
		logic.GetUser(resp, ctx)
	case "/test":
		resp.SetValue("test")
	default:
		resp.SetError(err.GetErr(2, ""))
	}
}
