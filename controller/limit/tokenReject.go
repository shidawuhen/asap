package limit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// @Tags limit
// @Summary 令牌桶拒流
// @Produce  json
// @Success 200 {string} string "成功会返回ok"
// @Failure 502 "失败返回reject"
// @Router /limit/tokenreject [get]
type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = 0
	l.lastTokenSec = time.Now().Unix()
}

func CreateTokenBucket() *TokenBucket {
	t := &TokenBucket{}
	t.Set(1, 5)
	return t
}

var tokenBucket *TokenBucket = CreateTokenBucket()

func TokenReject(c *gin.Context) {
	//fmt.Println(tokenBucket.tokens)
	if !tokenBucket.Allow() {
		c.String(http.StatusBadGateway, "reject")
		return
	}
	c.String(http.StatusOK, "ok")
}
