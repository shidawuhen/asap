/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package valueobject

type WarehouseId struct {
	warehouseId int64
}

func (w *WarehouseId) Get() int64 {
	return w.warehouseId
}

func NewWarehouseId(warehouseId int64) *WarehouseId {
	return &WarehouseId{
		warehouseId: warehouseId,
	}
}
