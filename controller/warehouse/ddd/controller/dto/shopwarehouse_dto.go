/**
@author: Jason Pang
@desc:
@date: 2022/7/21
**/
package dto

type ShopWarehouseCreateDTO struct {
	Code          string
	Name          string
	SpWarehouseId int64
}

type ShopWarehouseUpdateStatusDTO struct {
	ShopWarehouseId int64
	Status          int64
}
