package main

import (
	"asap/aredis"
	"asap/controller/algorithm"
	_ "asap/docs"
	"asap/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//nums := [][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}
	//nums := [][]int{{1, 2, 7}, {3, 6, 7}}
	//nums := [][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}
	nums := [][]int{{1, 0},{1,2},{0,1}}
	fmt.Println(algorithm.CanFinish(3,nums))
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
