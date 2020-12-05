package limit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

/*
// 定义漏桶结构
type LeakyBucket struct {
	capacity int64  // 桶的容量（接受缓存的请求总量）
	rate  int64// 水流出的速度（处理请求速度）
	water int64 // 当前水量（当前累计请求数）
	lastTokenSec int64 // 当前注水时间戳 （当前请求时间戳）
	lock sync.Mutex
}

// 判断是否加水（是否处理请求）
func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	// 先执行漏水，计算剩余水量
	l.water = l.water - (now-l.lastTokenSec)*l.rate
	if l.water < 0 {
		l.water = 0
	}
	l.lastTokenSec = now

	if l.water < l.capacity {
		// 尝试加水，此时水桶未满
		l.water++
		return true
	}else {
		// 水满了，拒绝加水
		return false
	}
}

func (l *LeakyBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.lastTokenSec = time.Now().Unix()
}

func CreateLeakyBucket() *LeakyBucket {
	t := &LeakyBucket{}
	t.Set(1, 5)
	return t
}

type rateLimiter struct {
	lck      *sync.Mutex
	rate     float64   //最大速率限制
	balance  float64   //漏桶的余量
	limit    float64   //漏桶的最大容量限制
	lastTime time.Time //上次检查的时间
}

func NewRateLimiter(limitPerSecond int, balance int) *rateLimiter {
	return &rateLimiter{
		lck:      new(sync.Mutex),
		rate:     float64(limitPerSecond),
		balance:  float64(balance),
		limit:    float64(balance),
		lastTime: time.Now(),
	}
}

func (r *rateLimiter) Check() bool {
	ok := false
	r.lck.Lock()
	now := time.Now()
	dur := now.Sub(r.lastTime).Seconds()
	r.lastTime = now
	water := dur * r.rate //计算这段时间内漏桶流出水的流量water
	r.balance += water    //漏桶流出water容量的水，自然漏桶的余量多出water
	if r.balance > r.limit {
		r.balance = r.limit
	}
	fmt.Println(r.balance)
	if r.balance >= 1 { //漏桶余量足够容下当前的请求
		r.balance -= 1
		ok = true
	}
	r.lck.Unlock()
	return ok
}

var leakyBucket *LeakyBucket = CreateLeakyBucket()
*/

//计数拒流变种
type LeakyBucket struct {
	// 容量
	capacity  int64
	// 剩余大小
	remaining int64
	// 下一次的重置容量时间
	reset     time.Time
	// 重置容量时间间隔
	rate      time.Duration
	mutex     sync.Mutex
}
func (b *LeakyBucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if time.Now().After(b.reset) { // 需要重置
		b.reset = time.Now().Add(b.rate) // 更新时间
		b.remaining = b.capacity // 重置剩余容量
	}
	fmt.Println(b.remaining)
	if b.remaining > 0 { // 判断是否能过
		b.remaining--
		return true
	}
	return false
}

func (b *LeakyBucket) Set(r time.Duration, c int64) {
	b.rate = r
	b.capacity = c
	b.remaining = c
	b.reset = time.Now().Add(b.rate)
}

func CreateLeakyBucket(r time.Duration,c int64) *LeakyBucket {
	t := &LeakyBucket{}
	t.Set(r, c)
	return t
}

var leakyBucket *LeakyBucket = CreateLeakyBucket(time.Second*2,10)
func LeakyReject(c *gin.Context) {
	if !leakyBucket.Allow() {
		c.String(http.StatusBadGateway, "reject")
		return
	}
	c.String(http.StatusOK, "ok")
}

