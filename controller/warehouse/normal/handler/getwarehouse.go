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

func GetWareHouse(c *gin.Context) {
	//参数检查
	//获取信息返回
	warehouseService := service.NewShopWareHouseService()
	warehouseInfo := warehouseService.GetShopWareHouse(1)
	warehouseInfoStr, _ := json.Marshal(warehouseInfo)
	c.String(http.StatusOK, string(warehouseInfoStr))
}
