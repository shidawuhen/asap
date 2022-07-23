/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package po

import "time"

type ShopWareHouse struct {
	Id              int64     `gorm:"id" json:"id"`
	WarehouseId     int64     `gorm:"warehouse_id" json:"warehouse_id"`
	Code            string    `gorm:"code" json:"code"`
	Name            string    `gorm:"name" json:"name"`
	SpWareHouseId   int64     `gorm:"column:scsp_warehouse_id" json:"scsp_warehouse_id"`
	SpWareHouseCode string    `gorm:"column:scsp_warehouse_code" json:"scsp_warehouse_code"`
	Status          int64     `gorm:"column:Status" json:"Status"`
	Updater         string    `gorm:"column:update_by" json:"update_by"`
	UpdateTime      time.Time `gorm:"update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"create_time" json:"create_time"`
}
