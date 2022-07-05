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

/**
 * @Author: Jason Pang
 * @Description: 获取商家仓信息
 * @receiver s
 * @param id
 * @return *model.ShopWareHouse
 */
func (s *ShopWareHouseRepo) GetShopWareHouse(id int64) *model.ShopWareHouse {
	return &model.ShopWareHouse{
		Id:            1,
		WarehouseId:   11,
		Code:          "商家仓1",
		Name:          "商家仓1",
		SpWareHouseId: 2,
	}
}

/**
 * @Author: Jason Pang
 * @Description: 创建商家仓
 * @receiver s
 * @param info
 * @return bool
 */
func (s *ShopWareHouseRepo) CreateShopWareHouse(info *model.ShopWareHouse) bool {
	return true
}
