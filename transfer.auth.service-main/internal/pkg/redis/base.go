package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	client *redis.Client
	host   string
)

const (
	defaultTimeout = 30
	defaultTtl     = time.Second * 3600
)

func Init(dbhost string) error {
	host = dbhost
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	cmd := client.Ping(context.Background())
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func Set(key string, val interface{}, ttl ...time.Duration) error {
	var ttlVal time.Duration
	if len(ttl) > 0 {
		ttlVal = ttl[0]
	} else {
		ttlVal = defaultTtl
	}
	err := client.Set(getContext(defaultTimeout), key, val, ttlVal).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string, in interface{}) error {
	var err error
	in, err = client.Get(getContext(defaultTimeout), key).Result()
	if err != nil {
		return err
	}
	return nil //returns redis.Nil if nothing found
}

func Keys(key string, val interface{}) error {
	cmd := client.Keys(getContext(defaultTimeout), key)
	if err := cmd.Err(); err != nil {
		return err
	}
	var list []string
	if err := cmd.ScanSlice(&list); err != nil {
		return err
	}
	if len(list) > 0 {
		cmdList := client.MGet(getContext(defaultTimeout), list...)
		if err := cmdList.Err(); err != nil {
			return err
		} else {
			res := cmdList.Val()

			if len(res) == 0 {
				return nil
			}
			if len(res) > 0 {
				var ins []interface{}
				for _, el := range res {
					var in interface{}
					if err := json.Unmarshal([]byte(el.(string)), &in); err == nil {
						ins = append(ins, in)
					}
				}
				if len(ins) > 0 {
					bt, err := json.Marshal(ins)
					if err != nil {
						return err
					}
					if err := json.Unmarshal(bt, val); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// not finish
func KeysRightSearch(key string, val interface{}) error {
	var err error
	val, _, err = client.Scan(getContext(defaultTimeout), 0, key+"*", 0).Result()
	if err != nil {
		return err
	}
	return nil
}

// not finish
func KeysLeftSearch(key string, val interface{}) error {
	var err error
	val, _, err = client.Scan(getContext(defaultTimeout), 0, "*"+key, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetPipe() redis.Pipeliner {
	return client.Pipeline()
}

func getContext(w int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(w)*time.Second)
	return ctx
}
