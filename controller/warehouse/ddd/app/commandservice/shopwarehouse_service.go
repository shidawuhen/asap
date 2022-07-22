/**
@author: Jason Pang
@desc: 应用层，直接就是handle调用了，但是感觉没必要使用接口类型。
我即使用的是实体，后面也可以按照函数建个新的，只是需要改一下每一个调用位置
@date: 2022/7/8
**/
package commandservice

import (
	"asap/controller/warehouse/ddd/domain/command"
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/domain/repo"
	"asap/controller/warehouse/ddd/integration/acl"
	"context"
	"errors"
)

type ShopWarehouseApplicationService struct {
	ShopWarehouseRepo repo.ShopWarehouseRepository //这个是个接口类型
	SpWarehouseRepo   repo.SpWarehouseRepository
	WarehouseAcl      *acl.WarehouseAcl
	ctx               context.Context
}

func NewShopWarehouseApplicationService(
	ctx context.Context,
	shopWarehouseRepo repo.ShopWarehouseRepository,
	spWarehouseRepo repo.SpWarehouseRepository,
	warehouseAcl *acl.WarehouseAcl,
) *ShopWarehouseApplicationService {
	return &ShopWarehouseApplicationService{
		ctx:               ctx,
		ShopWarehouseRepo: shopWarehouseRepo,
		SpWarehouseRepo:   spWarehouseRepo,
		WarehouseAcl:      warehouseAcl,
	}
}

// 这里先暂时忽略服务方法的入参、出参
func (s *ShopWarehouseApplicationService) Create(command *command.ShopWarehouseCreateCommand) error {
	//1.从数据库获取数据，补充服务商仓
	s.SpWarehouseRepo.Find(s.ctx, command.SpWarehouse.WarehouseId)
	//2.通过rpc，补充warehouseid
	warehouseId := s.WarehouseAcl.GetWarehouseId(command.Code)
	command.WarehouseId = *warehouseId
	//3.调用聚合创建
	shopWarehouseAggregate := aggregate.ShopWarehouse{}
	shopWarehouse := shopWarehouseAggregate.Create(command)
	//4.存储
	s.ShopWarehouseRepo.Save(s.ctx, shopWarehouse)
	return nil
}

//update等
func (s *ShopWarehouseApplicationService) UpdateStatus(command *command.ShopWarehouseUpdateStatusCommand) error {
	//1.从数据库获取商家仓信息
	shopWareInfo, _ := s.ShopWarehouseRepo.Find(s.ctx, command.WarehouseId.Get())
	//2.调用聚合更新状态
	shopWarehouseAggregate := aggregate.ShopWarehouse{}
	shopWarehouse := shopWarehouseAggregate.UpdateStatus(command, shopWareInfo)
	if shopWarehouse == nil {
		return errors.New("更新失败")
	}
	//3.存储
	s.ShopWarehouseRepo.Save(s.ctx, shopWarehouse)
	return nil
}
