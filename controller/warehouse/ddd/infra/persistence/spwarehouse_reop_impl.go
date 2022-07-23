/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package persistence

import (
	"asap/controller/warehouse/ddd/domain/model/entity"
	"asap/controller/warehouse/ddd/infra/persistence/convertor"
	"asap/controller/warehouse/ddd/infra/persistence/dal"
	"context"
)

type SpWarehouseRepository struct {
}

func NewSpWarehouseRepository() *SpWarehouseRepository {
	return &SpWarehouseRepository{}
}

func (r *SpWarehouseRepository) Find(ctx context.Context, id int64) (*entity.SpWarehouse, error) {
	//操作db根据id获取到sp的信息，mock结果出来
	dal := dal.NewSpWarehouseDalRepository()
	spwareInfo := dal.Find(id)
	convertor := convertor.WarehouseConvertor{}
	return convertor.CreateSpWarehouseEntity(spwareInfo), nil
}
