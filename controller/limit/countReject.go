package limit

import (
	"asap/aredis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Tags limit
// @Summary 计数拒流，每秒超过指定次数会拒流掉
// @Produce  json
// @Success 200 {string} string "成功会返回ok"
// @Failure 502 "失败返回reject"
// @Router /limit/countreject [get]
func CountReject(c *gin.Context) {
	currentTime := time.Now().Unix()
	key := fmt.Sprintf("count:%d", currentTime)
	limitCount := 5
	fmt.Println(key)
	trafficCount, _ := aredis.GetRedis(aredis.BASEREDIS).Incr(key)
	if trafficCount == 1 {
		aredis.GetRedis(aredis.BASEREDIS).Expire(key, 86400)
	}
	if int(trafficCount) > limitCount {
		c.String(http.StatusBadGateway, "reject")
		return
	}
	c.String(http.StatusOK, "ok")
}
