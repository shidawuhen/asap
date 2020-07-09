package main

import (
	"asap/aredis"
	_ "asap/docs"
	"asap/router"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	initRedis()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}

func initRedis(){
	myRedis := newRedisManager("myRedis")
	aredis.SetRedis("myRedis",myRedis)
}

func newRedisManager(servicename string) (redis *aredis.RedisManager) {
	redis, err := aredis.NewRedisManager(servicename, "127.0.0.1:6379", "111111", 5, 2000, 2000, 2000)
	if err != nil {
		panic(err)
	}

	return
}