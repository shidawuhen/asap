/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package command

import "asap/controller/warehouse/ddd/domain/model/entity"

/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package command

import "asap/controller/warehouse/ddd/domain/model/entity"

type ShopWarehouseCreateCommand struct {
	WarehouseId int64
	Code        string
	Name        string
	SpWarehouse entity.SpWarehouse
}
