package limit

import (
	"fmt"
	"github.com/andres-erbsen/clock"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

//真固定速率
type Clock interface {
	Now() time.Time
	Sleep(time.Duration)
}

type limiter struct {
	sync.Mutex               // 锁
	last       time.Time     // 上一次的时刻
	sleepFor   time.Duration // 需要等待的时间
	perRequest time.Duration // 每次的时间间隔
	maxSlack   time.Duration // 最大的富余量
	clock      Clock         // 时钟
}

// Take 会阻塞确保两次请求之间的时间走完
// Take 调用平均数为 time.Second/rate.
func (t *limiter) Take() time.Time {
	t.Lock()
	defer t.Unlock()

	now := t.clock.Now()

	// 如果是第一次请求就直接放行
	if t.last.IsZero() {
		t.last = now
		return t.last
	}

	// sleepFor 根据 perRequest 和上一次请求的时刻计算应该sleep的时间
	// 由于每次请求间隔的时间可能会超过perRequest, 所以这个数字可能为负数，并在多个请求之间累加
	t.sleepFor += t.perRequest - now.Sub(t.last)
	fmt.Println(t.sleepFor)
	// 我们不应该让sleepFor负的太多，因为这意味着一个服务在短时间内慢了很多随后会得到更高的RPS。
	if t.sleepFor < t.maxSlack {
		t.sleepFor = t.maxSlack
	}

	// 如果 sleepFor 是正值那么就 sleep
	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
		t.last = now.Add(t.sleepFor)
		t.sleepFor = 0
	} else {
		t.last = now
	}
	return t.last
}

func NewLimiter(rate int) *limiter {
	l := &limiter{
		perRequest: time.Second / time.Duration(rate),
		maxSlack:   -10 * time.Second / time.Duration(rate),
	}

	if l.clock == nil {
		l.clock = clock.New()
	}
	return l
}

var rl = NewLimiter(100) // per second,每秒100个请求
func LeakyRejectFixedRate(c *gin.Context) {
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
	c.String(http.StatusOK, "ok")
}
