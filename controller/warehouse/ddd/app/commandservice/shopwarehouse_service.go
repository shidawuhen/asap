/**
@author: Jason Pang
@desc: 应用层，直接就是handle调用了，但是感觉没必要使用接口类型。
我即使用的是实体，后面也可以按照函数建个新的，只是需要改一下每一个调用位置
@date: 2022/7/8
**/
package commandservice

import (
	"asap/controller/warehouse/ddd/domain/command"
	"asap/controller/warehouse/ddd/domain/repo"
	"context"
)

type ShopWarehouseApplicationService struct {
	ShopWarehouseRepo repo.ShopWarehouseRepository //这个是个接口类型
	SpWarehouseRepo   repo.SpWarehouseRepository
	ctx               context.Context
}

func NewShopWarehouseApplicationService(ctx context.Context,
	shopWarehouseRepo repo.ShopWarehouseRepository,
	spWarehouseRepo repo.SpWarehouseRepository) *ShopWarehouseApplicationService {
	return &ShopWarehouseApplicationService{
		ctx:               ctx,
		ShopWarehouseRepo: shopWarehouseRepo,
		SpWarehouseRepo:   spWarehouseRepo,
	}
}

// 这里先暂时忽略服务方法的入参、出参
func (s *ShopWarehouseApplicationService) Create(command *command.ShopWarehouseCreateCommand) error {
	//补充服务商仓
	s.SpWarehouseRepo.Find(s.ctx, command.SpWarehouse.WarehouseId)
	//补充warehouseid

	//调用聚合创建

	//存储
	return nil
}
