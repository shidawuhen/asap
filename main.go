package main

import (
	"asap/aredis"
	_ "asap/docs"
	"asap/router"
	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	"asap/controller/algorithm"
	"time"
	"fmt"
	"context"
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
	initETCD()
	r := gin.Default()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}

func initETCD(){
	var (
		config clientv3.Config
		err error
		client *clientv3.Client
		kv clientv3.KV
		getResp *clientv3.GetResponse

	)
	//配置
	config = clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*5,
	}
	//连接 床见一个客户端
	if client,err = clientv3.New(config);err != nil{
		fmt.Println(err)
		return
	}

	serviceName := "/"+GROUP+ "/" + TEAM + "/"
	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)
	getResp,err = kv.Get(context.TODO(),serviceName,clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range getResp.Kvs{
		fmt.Println(string(v.Value))
		global.SetService(serviceName,string(v.Value))
	}

	fmt.Println(global.GetServiceArr())
	fmt.Println(getResp.Kvs)
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