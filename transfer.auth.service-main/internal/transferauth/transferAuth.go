package transferauth

import (
	"fmt"
	"github.com/transferMVP/transfer.webapp/internal/config"
	"github.com/transferMVP/transfer.webapp/internal/delivery/http"
	"github.com/transferMVP/transfer.webapp/internal/repository/pg"
	"github.com/transferMVP/transfer.webapp/internal/service/handlers"
	"log"
)

var gfsd = make(chan bool)

func Run() {
	fmt.Println("run")
	var errCh chan error
	// инициализируем глобальную переменную конфиг
	if err := config.Init(); err != nil {
		errCh <- err
		log.Fatal("error load config : ", err)
	}

	//проверка на коннект пула
	if err := pg.InitPool(); err != nil {
		log.Fatal("error connect to db : ", err)
	}

	//проверка на коннект redis
	//if err := redis2.Init(config.Config.Redis.Dsn); err != nil {
	//	log.Fatal("error connect to redis : ", err)
	//}
	//google.Init()
	go http.Serv(handlers.RoutingHttp)

	go http.ServDocs("8888")
	<-gfsd
}
