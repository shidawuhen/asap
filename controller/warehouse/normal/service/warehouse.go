/**
@author: Jason Pang
@desc:
@date: 2022/7/3
**/
package service

import (
	"asap/controller/warehouse/normal/dal/db"
	"asap/controller/warehouse/normal/idl"
	"asap/controller/warehouse/normal/model"
)

type shopWarehouseService struct {
}

func NewShopWareHouseService() *shopWarehouseService {
	return &shopWarehouseService{}
}

/**
 * @Author: Jason Pang
 * @Description: 获取商家仓
 * @receiver s
 * @param id
 */
func (s *shopWarehouseService) GetShopWareHouse(id int64) *idl.ShopWareHouseInfo {
	//获取商家仓信息
	shopWareHouseRepo := db.DefaultShopWareHouseRepo()
	shopWareHouse := shopWareHouseRepo.GetShopWareHouse(id)
	//获取服务商仓信息
	spWareHouseRepo := db.DefaultSpWareHouseRepo()
	spWareHouse := spWareHouseRepo.GetSpWareHouse(shopWareHouse.SpWareHouseId)
	//组装
	return s.PackageShopWareHouseData(shopWareHouse, spWareHouse)
}

func (s *shopWarehouseService) PackageShopWareHouseData(shopWareHouse *model.ShopWareHouse, spWareHouse *model.SpWareHouse) *idl.ShopWareHouseInfo {
	return &idl.ShopWareHouseInfo{
		Id:            shopWareHouse.Id,
		WarehouseId:   shopWareHouse.WarehouseId,
		Code:          shopWareHouse.Code,
		Name:          shopWareHouse.Name,
		SpWareHouseId: spWareHouse.Id,
		SpCode:        spWareHouse.Code,
		SpName:        spWareHouse.Name,
	}
}

/**
 * @Author: Jason Pang
 * @Description: 创建商家仓
 * @receiver s
 */
func (s *shopWarehouseService) CreateShopWareHouse() {

}
