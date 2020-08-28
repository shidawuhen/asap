package global

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"strings"
	"time"
	"fmt"
	"context"
)
var (
	config clientv3.Config
	err error
	client *clientv3.Client
	kv clientv3.KV
	getResp *clientv3.GetResponse

)
var (
	//
	globalService map[string](map[string]string)
)

func init() {
	globalService = make(map[string](map[string]string))
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
}

func SetService(serviceName string, address string) {
	if _, ok := globalService[serviceName];!ok {
		globalService[serviceName] = make(map[string]string)
	}
	globalService[serviceName][address] = address
}

func DelService(serviceName string, address string) bool{
	if _,ok:= globalService[serviceName];ok{
		if _,ok2 := globalService[serviceName][address];ok2{
			delete(globalService[serviceName],address)
			return true
		}
	}
	return false
}

func GetService(serviceName string) (map[string]string) {
	return globalService[serviceName]
}

func GetServiceArr() map[string](map[string]string) {
	return globalService
}

func GetServiceFromEtcd(serviceName string){
	if client == nil{
		return
	}
	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)
	getResp,err = kv.Get(context.TODO(),serviceName,clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range getResp.Kvs{
		fmt.Println(string(v.Value))
		SetService(serviceName,string(v.Value))
	}

	fmt.Println(GetServiceArr())
	fmt.Println(getResp.Kvs)
}

func WatchServiceFromEtcd(serviceName string){
	if client == nil{
		return
	}
	/*ticker := time.NewTicker(time.Second * 20)
	go func() {
		for range ticker.C {

		}
	}*/
	// 创建一个watcher
	watcher := clientv3.NewWatcher(client)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(500000 * time.Second, func() {
		cancelFunc()
	})

	watchRespChan := watcher.Watch(ctx, serviceName, clientv3.WithPrefix())

	// 处理kv变化事件
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
				SetService(serviceName,string(event.Kv.Value))
				fmt.Println("now service ip", GetService(serviceName))
			case mvccpb.DELETE:
				fmt.Println("删除了" + strings.TrimPrefix(string(event.Kv.Key),serviceName), "Revision:", event.Kv.ModRevision)
				DelService(serviceName,strings.TrimPrefix(string(event.Kv.Key),serviceName))
				fmt.Println("now service ip", GetService(serviceName))
			}
		}
	}
}


