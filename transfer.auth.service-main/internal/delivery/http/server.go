package http

import (
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/transferMVP/transfer.webapp/docs"
	"github.com/transferMVP/transfer.webapp/internal/config"
	"github.com/valyala/fasthttp"
	"net/http"
)

func Serv(proc func(ctx *fasthttp.RequestCtx)) {
	fasthttp.ListenAndServe(config.Config.Server.Addr, proc).Error()
}

func ServDocs(port string) {
	fmt.Println(port)
	router := mux.NewRouter()

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	if err := http.ListenAndServe("0.0.0.0:"+port, router); err != nil {
		fmt.Println(err)
	}
}
