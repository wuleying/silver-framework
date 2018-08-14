package redis

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/go-redis/redis"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"strconv"
)

type Redis struct {
	Config config.Config
}

// Init 初始化redis
func (r *Redis) Init() {
	db, err := strconv.Atoi(r.Config["redis"]["db"])
	exceptions.CheckError(err)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Config["redis"]["host"], r.Config["redis"]["port"]),
		Password: r.Config["redis"]["password"],
		DB:       db,
	})

	var counterKey = "counter"

	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()

	exceptions.CheckError(err)

	clog.Info("Current counter is %s", cntValue)
}
