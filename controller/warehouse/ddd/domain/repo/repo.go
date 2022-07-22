/**
@author: Jason Pang
@desc:
@date: 2022/7/8
**/
package repo

import (
	"asap/controller/warehouse/normal/model"
	"context"
)

type ShopWarehouseRepository interface {
	Save(context.Context, *model.ShopWareHouse) error          // 保存一个聚合
	Find(context.Context, int64) (*model.ShopWareHouse, error) // 通过id查找对应的聚合
}

type SpWarehouseRepository interface {
	Find(context.Context, int64) (*model.SpWareHouse, error) // 通过warehouseid查找服务商仓
}
