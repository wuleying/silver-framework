package redis

import (
	"github.com/go-clog/clog"
	"github.com/go-redis/redis"
	"github.com/wuleying/silver-framework/exceptions"
)

// Init 初始化redis
func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var counterKey = "counter"

	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()

	exceptions.CheckError(err)

	clog.Info("current counter is %s", cntValue)
}
