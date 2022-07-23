/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package acl

import "asap/controller/warehouse/ddd/domain/model/valueobject"

type WarehouseAcl struct {
}

func NewWarehouseAcl() *WarehouseAcl {
	return &WarehouseAcl{}
}

func (w *WarehouseAcl) GetWarehouseId(code string) *valueobject.WarehouseId {
	//通过rpc请求，使用code获取到warehosueid
	//rpc.GetWarehouseId(code)
	var warehouseId int64 = 111
	return valueobject.NewWarehouseId(warehouseId)
}
