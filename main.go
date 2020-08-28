package main

import (
	"asap/aredis"
	_ "asap/docs"
	"asap/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"asap/controller/algorithm"
	"asap/global"
)

const (
	GROUP = "b2c"
	TEAM =  "i18n"
)


func main() {
	nums := []int{1,2,3,4}
	algorithm.Exchange(nums)
	InitRedis()
	serviceName := "/"+GROUP+ "/" + TEAM + "/"
	global.GetServiceFromEtcd(serviceName)
	go global.WatchServiceFromEtcd(serviceName)
	fmt.Println(2)
	r := gin.Default()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8082
	r.Run(":8082")
}


func InitRedis(){
	myRedis := newRedisManager(aredis.BASEREDIS)
	aredis.SetRedis(aredis.BASEREDIS,myRedis)
}

func newRedisManager(servicename string) (redis *aredis.RedisManager) {
	redis, err := aredis.NewRedisManager(servicename, "127.0.0.1:6379", "111111", 5, 2000, 2000, 2000)
	if err != nil {
		panic(err)
	}

	return
}