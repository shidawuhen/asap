/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package assembler

import (
	"asap/controller/warehouse/ddd/controller/querydto"
	"asap/controller/warehouse/ddd/domain/command"
	"asap/controller/warehouse/ddd/domain/model/entity"
)

type Assembler struct {
}

func (a Assembler) ToCommandFromDTO(dto querydto.ShopWarehouseCreateDTO) *command.ShopWarehouseCreateCommand {
	return &command.ShopWarehouseCreateCommand{
		Code: dto.Code,
		Name: dto.Name,
		SpWarehouse: entity.SpWarehouse{
			WarehouseId: dto.SpWarehouseId,
		},
	}
}
