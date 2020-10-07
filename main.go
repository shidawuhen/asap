package main

import (
	"asap/aredis"
	"asap/controller/algorithm"
	_ "asap/docs"
	"asap/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//nums := []int{2}
	//edges := [][]int{{0,1},{0,2},{2,3},{2,4},{2,5}}
	speed := []int{9,9,8,5}
	efficiency := []int{9,2,1,6}
	k := 3
	algorithm.MaxPerformance(4,speed,efficiency,k)
	r := gin.Default()
	InitRedis()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}

func InitRedis() {
	myRedis := newRedisManager(aredis.BASEREDIS)
	aredis.SetRedis(aredis.BASEREDIS, myRedis)
}

func newRedisManager(servicename string) (redis *aredis.RedisManager) {
	redis, err := aredis.NewRedisManager(servicename, "127.0.0.1:6379", "111111", 5, 2000, 2000, 2000)
	if err != nil {
		panic(err)
	}

	return
}
