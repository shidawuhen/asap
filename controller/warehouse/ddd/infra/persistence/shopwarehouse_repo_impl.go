/**
@author: Jason Pang
@desc:
@date: 2022/7/8
**/
package persistence

import (
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/infra/persistence/convertor"
	"asap/controller/warehouse/ddd/infra/persistence/dal"
	"context"
)

type ShopWarehouseRepository struct {
}

func NewShopWarehouseRepository() *ShopWarehouseRepository {
	return &ShopWarehouseRepository{}
}

func (r *ShopWarehouseRepository) Save(ctx context.Context, ware *aggregate.ShopWarehouse) error {
	//将数据进行存储
	convert := convertor.WarehouseConvertor{}
	po := convert.CreateShopWarehousePO(ware)
	dal := dal.NewShopWarehouseRepository()
	return dal.Create(po)
}

func (r *ShopWarehouseRepository) Find(ctx context.Context, id int64) (*aggregate.ShopWarehouse, error) {
	dal := dal.NewShopWarehouseRepository()
	shopWareInfo := dal.Find(id)
	convert := convertor.WarehouseConvertor{}
	return convert.CreateShopWarehouseEntity(shopWareInfo), nil
}
