/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package queryservice

import (
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/domain/repo"
	"context"
)

type ShopWarehouseQueryApplicationService struct {
	ctx               context.Context
	ShopWarehouseRepo repo.ShopWarehouseRepository //这个是个接口类型
	SpWarehouseRepo   repo.SpWarehouseRepository
}

func NewShopWarehouseQueryApplicationService(
	ctx context.Context,
	shopWarehouseRepo repo.ShopWarehouseRepository,
	spWarehouseRepo repo.SpWarehouseRepository,
) *ShopWarehouseQueryApplicationService {
	return &ShopWarehouseQueryApplicationService{
		ctx:               ctx,
		ShopWarehouseRepo: shopWarehouseRepo,
		SpWarehouseRepo:   spWarehouseRepo,
	}
}

func (s *ShopWarehouseQueryApplicationService) GetShopWarehouse(warehouseId int64) *aggregate.ShopWarehouse {
	shopWarehouseInfo, _ := s.ShopWarehouseRepo.Find(s.ctx, warehouseId)
	return shopWarehouseInfo
}
