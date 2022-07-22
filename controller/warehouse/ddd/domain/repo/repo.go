/**
@author: Jason Pang
@desc:
@date: 2022/7/8
**/
package repo

import (
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/domain/model/entity"
	"context"
)

type ShopWarehouseRepository interface {
	Save(context.Context, *aggregate.ShopWarehouse) error          // 保存一个聚合，存在则更新，不存在则插入
	Find(context.Context, int64) (*aggregate.ShopWarehouse, error) // 通过id查找对应的聚合
}

type SpWarehouseRepository interface {
	Find(context.Context, int64) (*entity.SpWarehouse, error) // 通过warehouseid查找服务商仓
}
