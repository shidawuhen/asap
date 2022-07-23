/**
@author: Jason Pang
@desc: 处理服务商仓表
@date: 2022/7/3
**/
package db

import (
	"asap/controller/warehouse/normal/model"
	"sync"
)

type SpWareHouseRepo struct {
}

func NewSpWareHouseRepo() *SpWareHouseRepo {
	return &SpWareHouseRepo{}
}

var (
	defaultSpRepoOnce sync.Once
	defaultSpRepo     *SpWareHouseRepo
)

func DefaultSpWareHouseRepo() *SpWareHouseRepo {
	defaultSpRepoOnce.Do(func() {
		defaultSpRepo = NewSpWareHouseRepo()
	})
	return defaultSpRepo
}

/**
 * @Author: Jason Pang
 * @Description: 获取服务商仓
 * @receiver s
 * @param id
 * @return *model.SpWareHouse
 */
func (s *SpWareHouseRepo) GetSpWareHouse(id int64) *model.SpWareHouse {
	return &model.SpWareHouse{
		Id:   2,
		Code: "服务商仓2",
		Name: "服务商仓2",
	}
}
