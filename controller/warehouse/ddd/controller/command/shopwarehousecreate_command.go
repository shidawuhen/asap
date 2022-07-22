/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package command

import (
	"asap/controller/warehouse/ddd/domain/model/entity"
	"asap/controller/warehouse/ddd/domain/model/valueobject"
)

type ShopWarehouseCreateCommand struct {
	WarehouseId valueobject.WarehouseId
	Code        string
	Name        string
	SpWarehouse entity.SpWarehouse
}
