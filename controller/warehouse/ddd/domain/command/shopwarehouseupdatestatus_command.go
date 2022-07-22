/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package command

import (
	"asap/controller/warehouse/ddd/domain/model/valueobject"
)

type ShopWarehouseUpdateStatusCommand struct {
	WarehouseId *valueobject.WarehouseId
	Status      int64
}
