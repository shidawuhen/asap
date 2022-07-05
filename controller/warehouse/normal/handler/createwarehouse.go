/**
@author: Jason Pang
@desc:
@date: 2022/7/3
**/
package handler

import (
	"asap/controller/warehouse/normal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatWareHouse(c *gin.Context) {
	warehouseService := service.NewShopWareHouseService()
	res := warehouseService.CreateShopWareHouse()
	resStr, _ := json.Marshal(res)
	c.String(http.StatusOK, string(resStr))
}
