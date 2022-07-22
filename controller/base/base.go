package base

import (
	"asap/controller/warehouse/ddd/controller"
	"asap/controller/warehouse/ddd/controller/dto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags base
// @Summary 接口探活
// @Description ping接口，用于测试是否调通
// @Produce  json
// @Param name query string false "参数描述"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	controller := controller.NewShopWarehouseController(context.Background())
	updateDto := dto.ShopWarehouseUpdateStatusDTO{
		ShopWarehouseId: 1,
		Status:          1,
	}
	controller.UpdateStatus(updateDto)

	createDto := dto.ShopWarehouseCreateDTO{
		Code:          "商家仓1",
		Name:          "商家仓1",
		SpWarehouseId: 1,
	}
	controller.Create(createDto)

	controller.GetShopWarehouse(1)
	c.String(http.StatusOK, "ok")

}
