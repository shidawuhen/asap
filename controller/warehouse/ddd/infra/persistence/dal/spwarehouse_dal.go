/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package dal

import (
	"asap/controller/warehouse/ddd/infra/persistence/po"
	"time"
)

type SpWarehouseDal struct {
}

func NewSpWarehouseDalRepository() *SpWarehouseDal {
	return &SpWarehouseDal{}
}

func (dal *SpWarehouseDal) Find(id int64) *po.SpWareHouse {
	//操作db根据id获取到sp的信息，mock结果出来
	return &po.SpWareHouse{
		Id:         id,
		Code:       "服务商仓",
		Name:       "我是服务商仓",
		SpId:       1,
		UpdateTime: time.Now(),
	}

}
