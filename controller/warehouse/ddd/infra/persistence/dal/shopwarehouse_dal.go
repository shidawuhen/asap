/**
@author: Jason Pang
@desc:
@date: 2022/7/8
**/
package dal

import (
	"asap/controller/warehouse/ddd/infra/persistence/po"
	"time"
)

type ShopWarehouseDal struct {
}

func NewShopWarehouseRepository() *ShopWarehouseDal {
	return &ShopWarehouseDal{}
}

//真正的操作
func (dal *ShopWarehouseDal) Create(ware *po.ShopWareHouse) error {
	//操作db，插入数据
	return nil
}

func (dal *ShopWarehouseDal) Find(id int64) *po.ShopWareHouse {
	//操作db根据id获取到sp的信息，mock结果出来
	return &po.ShopWareHouse{
		Id:              id,
		WarehouseId:     1,
		Code:            "商家仓仓",
		Name:            "我是商家仓",
		SpWareHouseId:   1,
		SpWareHouseCode: "服务商仓",
		UpdateTime:      time.Now(),
	}
}
