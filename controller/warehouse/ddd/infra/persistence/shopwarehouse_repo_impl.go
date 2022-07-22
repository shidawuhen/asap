/**
@author: Jason Pang
@desc:
@date: 2022/7/8
**/
package persistence

import (
	"asap/controller/warehouse/normal/model"
	"context"
)

type ShopWarehouseRepository struct {
}

func NewShopWarehouseRepository() *ShopWarehouseRepository {
	return &ShopWarehouseRepository{}
}

func (r *ShopWarehouseRepository) Save(ctx context.Context, order *model.ShopWareHouse) error {
	return nil
}

func (r *ShopWarehouseRepository) Find(ctx context.Context, id int64) (*model.ShopWareHouse, error) {
	return nil, nil
}
