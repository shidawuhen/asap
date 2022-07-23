/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package convertor

import (
	"asap/controller/warehouse/ddd/domain/model/aggregate"
	"asap/controller/warehouse/ddd/domain/model/entity"
	"asap/controller/warehouse/ddd/domain/model/valueobject"
	"asap/controller/warehouse/ddd/infra/persistence/po"
	"time"
)

type WarehouseConvertor struct {
}

func (w *WarehouseConvertor) CreateSpWarehouseEntity(p *po.SpWareHouse) *entity.SpWarehouse {
	return &entity.SpWarehouse{
		WarehouseId: p.Id,
		Code:        p.Code,
		Name:        p.Name,
	}
}

func (w *WarehouseConvertor) CreateShopWarehousePO(ware *aggregate.ShopWarehouse) *po.ShopWareHouse {
	return &po.ShopWareHouse{
		WarehouseId:     ware.WarehouseId.Get(),
		Code:            ware.Code,
		Name:            ware.Name,
		SpWareHouseId:   ware.SpWarehouse.WarehouseId,
		SpWareHouseCode: ware.SpWarehouse.Code,
		Status:          ware.Status,
		UpdateTime:      time.Now(),
		CreateTime:      time.Now(),
	}
}

func (w *WarehouseConvertor) CreateShopWarehouseEntity(ware *po.ShopWareHouse) *aggregate.ShopWarehouse {
	return &aggregate.ShopWarehouse{
		WarehouseId: *valueobject.NewWarehouseId(ware.WarehouseId),
		Code:        ware.Code,
		Name:        ware.Name,
		Status:      ware.Status,
		SpWarehouse: entity.SpWarehouse{
			WarehouseId: ware.SpWareHouseId,
			Code:        ware.SpWareHouseCode,
		},
	}
}
