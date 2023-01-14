package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/transferMVP/transfer.webapp/internal/config"
	"github.com/transferMVP/transfer.webapp/internal/pkg/redis"
	"time"
)

func Init(dsn string) error {
	return redis.Init(dsn)
}

func AddToken(token string) error {
	pipe := redis.GetPipe()
	defer pipe.Close()
	bts, err := json.Marshal(token)
	if err != nil {
		fmt.Println("bytes", err)
	}
	pipe.Set(getContext(30), config.Config.Redis.Key+":"+token, string(bts), time.Second*50)

	if _, err := pipe.Exec(getContext(30)); err != nil {
		fmt.Println("error writing tokens to redis: " + err.Error())
	}
	return nil
}

func getContext(w int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(w)*time.Second)
	return ctx
}
