package main

import (
	"asap/aredis"
	"asap/controller/algorithm"
	"asap/controller/design"
	_ "asap/docs"
	"asap/router"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/yedf/dtm/common"
	"github.com/yedf/dtm/dtmcli"
	"github.com/yedf/dtm/dtmcli/dtmimp"
)

// 启动命令：go run app/main.go qs

// 事务参与者的服务地址
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8082
const DtmServer = "http://localhost:8080/api/dtmsvr"

var qsBusi = fmt.Sprintf("http://localhost:%d%s", qsBusiPort, qsBusiAPI)

func qsAdjustBalance(uid int, amount int) (m interface{}, err error) {
	//_, err := dtmimp.DBExec(sdbGet(), "update dtm_busi.user_account set balance = balance + ? where user_id = ?", amount, uid)
	return dtmcli.MapSuccess, err
}

func qsAddRoute(app *gin.Engine) {
	app.POST(qsBusiAPI+"/TransIn", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
		return qsAdjustBalance(2, 30)
	}))
	app.POST(qsBusiAPI+"/TransInCompensate", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
		return qsAdjustBalance(2, -30)
	}))
	app.POST(qsBusiAPI+"/TransOut", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
		return qsAdjustBalance(1, -30)
	}))
	app.POST(qsBusiAPI+"/TransOutCompensate", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
		return qsAdjustBalance(1, 30)
	}))
}

// QsStartSvr 1
func QsStartSvr() {
	app := common.GetGinApp()
	qsAddRoute(app)
	dtmimp.Logf("quick qs examples listening at %d", qsBusiPort)
	go app.Run(fmt.Sprintf(":%d", qsBusiPort))
	time.Sleep(100 * time.Millisecond)
}

// QsFireRequest 1
func QsFireRequest() string {
	req := &gin.H{"amount": 30} // 微服务的载荷
	// DtmServer为DTM服务的地址
	saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
		// 添加一个TransOut的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransOutCompensate"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCompensate", req).
		// 添加一个TransIn的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransInCompensate"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCompensate", req)
	// 提交saga事务，dtm会完成所有的子事务/回滚所有的子事务
	err := saga.Submit()
	dtmimp.FatalIfError(err)
	return saga.Gid
}


func main() {
	single := design.GetSingleInstance()
	single.Show()
	nums := []int{3, 2, 1, 4, 5, 0}
	//nums := [][]byte{{'a'}}
	res := algorithm.GetLeastNumbers(nums, 2)
	fmt.Println(res)
	r := gin.Default()
	pprof.Register(r)
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
