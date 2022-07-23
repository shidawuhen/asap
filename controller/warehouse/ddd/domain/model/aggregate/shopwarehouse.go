/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package aggregate

import (
	"asap/controller/warehouse/ddd/domain/command"
	"asap/controller/warehouse/ddd/domain/model/entity"
	"asap/controller/warehouse/ddd/domain/model/valueobject"
)

type ShopWarehouse struct {
	WarehouseId valueobject.WarehouseId
	Code        string
	Name        string
	Status      int64
	SpWarehouse entity.SpWarehouse
}

/**
 * @Author: Jason Pang
 * @Description: 创建商家仓
 * @receiver s
 * @param command
 * @return *ShopWarehouse
 */
func (s *ShopWarehouse) Create(command *command.ShopWarehouseCreateCommand) *ShopWarehouse {
	return &ShopWarehouse{
		WarehouseId: command.WarehouseId,
		Code:        command.Code,
		Name:        command.Name,
		Status:      valueobject.SHOPWAREHOUSESTATUS_INIT,
		SpWarehouse: command.SpWarehouse,
	}
}

func (s *ShopWarehouse) UpdateStatus(command *command.ShopWarehouseUpdateStatusCommand, shopWare *ShopWarehouse) *ShopWarehouse {
	//此处是核心逻辑，判断更新的标准
	if shopWare.Status != command.Status {
		return nil
	}
	shopWare.Status = command.Status
	return shopWare
}
