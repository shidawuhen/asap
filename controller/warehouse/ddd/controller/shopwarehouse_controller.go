/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package controller

import (
	"asap/controller/warehouse/ddd/app/commandservice"
	"asap/controller/warehouse/ddd/app/queryservice"
	"asap/controller/warehouse/ddd/controller/assembler"
	"asap/controller/warehouse/ddd/controller/dto"
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/infra/persistence"
	"asap/controller/warehouse/ddd/integration/acl"
	"context"
	"fmt"
)

type shopWarehouseController struct {
	ctx            context.Context
	commandService *commandservice.ShopWarehouseApplicationService
	queryService   *queryservice.ShopWarehouseQueryApplicationService
}

func NewShopWarehouseController(ctx context.Context) *shopWarehouseController {
	return &shopWarehouseController{
		ctx: ctx,
		commandService: commandservice.NewShopWarehouseApplicationService(
			ctx,
			persistence.NewShopWarehouseRepository(),
			persistence.NewSpWarehouseRepository(),
			acl.NewWarehouseAcl(),
		),
		queryService: queryservice.NewShopWarehouseQueryApplicationService(
			ctx,
			persistence.NewShopWarehouseRepository(),
			persistence.NewSpWarehouseRepository(),
		),
	}
}
func (s *shopWarehouseController) Create(dto dto.ShopWarehouseCreateDTO) (err error) {
	assembler := assembler.NewAssembler()
	command := assembler.ToCommandFromCreateDTO(dto)
	fmt.Println("创建")
	return s.commandService.Create(command)
}

func (s *shopWarehouseController) UpdateStatus(dto dto.ShopWarehouseUpdateStatusDTO) (err error) {
	assembler := assembler.NewAssembler()
	command := assembler.ToCommandFromUpdateStatusDTO(dto)
	fmt.Println("更新")
	return s.commandService.UpdateStatus(command)
}

func (s *shopWarehouseController) GetShopWarehouse(warehouseId int64) *aggregate.ShopWarehouse {
	info := s.queryService.GetShopWarehouse(warehouseId)
	fmt.Println("查询", info)
	return info
}
