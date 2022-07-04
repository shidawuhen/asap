/**
@author: Jason Pang
@desc: 处理商家仓表
@date: 2022/7/3
**/
package db

import (
	"asap/controller/warehouse/normal/model"
	"sync"
)

type ShopWareHouseRepo struct {
}

func NewShopWareHouseRepo() *ShopWareHouseRepo {
	return &ShopWareHouseRepo{}
}

var (
	defaultRepoOnce sync.Once
	defaultRepo     *ShopWareHouseRepo
)

func DefaultShopWareHouseRepo() *ShopWareHouseRepo {
	defaultRepoOnce.Do(func() {
		defaultRepo = NewShopWareHouseRepo()
	})
	return defaultRepo
}

func (s *ShopWareHouseRepo) GetShopWareHouse(id int64) *model.ShopWareHouse {
	return &model.ShopWareHouse{
		Id:            1,
		WarehouseId:   11,
		Code:          "商家仓1",
		Name:          "商家仓1",
		SpWareHouseId: 2,
	}
}
