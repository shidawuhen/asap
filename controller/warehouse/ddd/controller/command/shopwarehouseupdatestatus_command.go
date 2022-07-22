/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package command

import (
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/domain/model/valueobject"
)

type ShopWarehouseUpdateStatusCommand struct {
	WarehouseId   valueobject.WarehouseId
	ShopWarehouse aggregate.ShopWarehouse
	Status        int64
}
