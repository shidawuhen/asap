package main

import (
	"asap/aredis"
	"asap/controller/algorithm"
	_ "asap/docs"
	"asap/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"asap/global"
)

const (
	GROUP = "b2c"
	TEAM =  "i18n"
)

func main() {
	//nums := [][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}
	//nums := [][]int{{1, 2, 7}, {3, 6, 7}}
	//nums := [][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}
	nums := [][]int{{2, 8}, {2}}
	algorithm.NumBusesToDestination(nums,8,2)
	r := gin.Default()
	InitRedis()
	serviceName := "/"+GROUP+ "/" + TEAM + "/"
	global.GetServiceFromEtcd(serviceName)
	go global.WatchServiceFromEtcd(serviceName)
	fmt.Println(2)
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8082
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
