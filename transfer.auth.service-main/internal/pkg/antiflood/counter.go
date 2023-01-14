package antiflood

import (
	"errors"
	"github.com/valyala/fasthttp"
	"net"
	"sync"
	"time"
)

var (
	num int = 10
	per int = 1

	clients = make(map[string]*ClientReq)
	mtx     sync.Mutex
)

type ClientReq struct {
	Count int
	Time  time.Time
}

// InitFreq takes 2 arguments: number of occurrences and period (in sec). Default: (10, 1sec)
func InitFreq(number, period int) {
	num = number
	per = period
}

// FastHttpCounter takes a request context of FastHTTP package, returns error.
// If error == nil, a request should be handled.
func FastHttpCounter(ctx *fasthttp.RequestCtx) error {
	err := Counter(ctx.RemoteIP(), ctx.Request.Header.Peek("User-Agent"))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusTooManyRequests)
		return err
	}
	return nil
}

// Counter takes IP and UserAgent info from request, returns error.
// If error == nil, a request should be handled.
func Counter(IP net.IP, UserAgent []byte) error {

	k := IP.String() + " " + string(UserAgent)

	mtx.Lock()
	defer mtx.Unlock()

	_, ok := clients[k]

	if !ok {
		clients[k] = &ClientReq{
			Count: 1,
			Time:  time.Now(),
		}
		return nil
	}

	if time.Duration(per)*time.Second > time.Now().Sub(clients[k].Time) {
		if clients[k].Count <= num {

			clients[k].Count++
			return nil
		}
		return errors.New("too many requests")
	}

	clients[k] = &ClientReq{
		Count: 1,
		Time:  time.Now(),
	}

	return nil
}
