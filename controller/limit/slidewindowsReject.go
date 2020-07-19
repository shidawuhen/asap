package limit

import (
	"container/ring"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
	"fmt"
)

var (
	limitCount int = 20 // 1s限频
	limitBucket int = 10 // 滑动窗口个数
	curCount int32 = 0  // 记录限频数量
	head *ring.Ring     // 环形队列（链表）
	printRes = 0
)

func init(){
	// 初始化滑动窗口
	head = ring.New(limitBucket)
	for i := 0; i < limitBucket; i++ {
		head.Value = 0
		head = head.Next()
	}
	// 启动执行器
	go func() {
		//ms级别，limitBucket int = 10意味将每秒分为10份，每份100ms
		timer := time.NewTicker(time.Millisecond * time.Duration(1000/limitBucket))
		for range timer.C { // 定时每隔指定时间刷新一次滑动窗口数据
			//subCount的作用，是因为当移动到head的时候，意味着该head要被废弃了。所以总count的值需要减去
			//head的值，并将head的值重新赋值为0
			subCount := int32(0 - head.Value.(int))
			newCount := atomic.AddInt32(&curCount, subCount)

			arr := make([]int,limitBucket)
			for i := 0; i < limitBucket; i++ { //打印出当前每个窗口的请求数量
				arr[i] = head.Value.(int)
				head = head.Next()
			}
			if printRes == 1 {
				fmt.Println("move subCount,newCount,arr", subCount, newCount,arr)
			}
			head.Value = 0
			head = head.Next()
		}
	}()
}
// @Tags limit
// @Summary 滑动窗口计数拒流，每秒超过指定次数会拒流掉
// @Produce  json
// @Success 200 {string} string "成功会返回ok"
// @Failure 502 "失败返回reject"
// @Router /limit/slidewindowsreject [get]
func SlideWindowsReject(c *gin.Context){
	n := atomic.AddInt32(&curCount, 1)
	if n > int32(limitCount) { // 超出限频
		atomic.AddInt32(&curCount, -1) //将多增加的数据减少
		c.String(http.StatusBadGateway, "reject")
	} else {
		mu := sync.Mutex{}
		mu.Lock()
		pos := head.Prev()
		val := pos.Value.(int)
		val++
		pos.Value = val
		mu.Unlock()
		c.String(http.StatusOK, "ok")
	}
}