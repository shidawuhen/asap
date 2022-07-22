/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package controller

import (
	"asap/controller/warehouse/ddd/app/commandservice"
	"asap/controller/warehouse/ddd/controller/assembler"
	"asap/controller/warehouse/ddd/controller/querydto"
	"asap/controller/warehouse/ddd/infra/persistence"
	"context"
)

type shopWarehouseController struct {
	ctx            context.Context
	commandService *commandservice.ShopWarehouseApplicationService
}

func NewShopWarehouseController(ctx context.Context) *shopWarehouseController {
	return &shopWarehouseController{
		ctx: ctx,
		commandService: commandservice.NewShopWarehouseApplicationService(
			ctx,
			persistence.NewShopWarehouseRepository(),
			persistence.NewSpWarehouseRepository()),
	}
}
func (s *shopWarehouseController) Create(dto querydto.ShopWarehouseCreateDTO) (err error) {
	command := assembler.Assembler{}.ToCommandFromDTO(dto)
	return s.commandService.Create(command)
}
