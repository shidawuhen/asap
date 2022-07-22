/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package assembler

import (
	"asap/controller/warehouse/ddd/controller/dto"
	"asap/controller/warehouse/ddd/domain/command"
	"asap/controller/warehouse/ddd/domain/model/entity"
	"asap/controller/warehouse/ddd/domain/model/valueobject"
)

type Assembler struct {
}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (a *Assembler) ToCommandFromCreateDTO(dto dto.ShopWarehouseCreateDTO) *command.ShopWarehouseCreateCommand {
	return &command.ShopWarehouseCreateCommand{
		Code: dto.Code,
		Name: dto.Name,
		SpWarehouse: entity.SpWarehouse{
			WarehouseId: dto.SpWarehouseId,
		},
	}
}

func (a *Assembler) ToCommandFromUpdateStatusDTO(dto dto.ShopWarehouseUpdateStatusDTO) *command.ShopWarehouseUpdateStatusCommand {
	return &command.ShopWarehouseUpdateStatusCommand{
		WarehouseId: *valueobject.NewWarehouseId(dto.ShopWarehouseId),
		Status:      dto.Status,
	}
}
