/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package persistence

import (
	"asap/controller/warehouse/normal/model"
	"context"
)

type SpWarehouseRepository struct {
}

func NewSpWarehouseRepository() *SpWarehouseRepository {
	return &SpWarehouseRepository{}
}

func (r *SpWarehouseRepository) Find(ctx context.Context, id int64) (*model.SpWareHouse, error) {
	return nil, nil
}
