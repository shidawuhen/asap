/**
@author: Jason Pang
@desc:
@date: 2022/7/3
**/
package model

/**
 * @Author: Jason Pang
 * @Description: 商家仓模型
 */
type ShopWareHouse struct {
	Id            int64  `gorm:"id" json:"id"`
	WarehouseId   int64  `gorm:"warehouse_id" json:"warehouse_id"`
	Code          string `gorm:"code" json:"code"`
	Name          string `gorm:"name" json:"name"`
	SpWareHouseId int64  `gorm:"column:scsp_warehouse_id" json:"scsp_warehouse_id"`
	Status        int64  `gorm:"status" json:"status"`
}

/**
 * @Author: Jason Pang
 * @Description: 服务商仓模型，服务商仓：商家仓=1：1
 */
type SpWareHouse struct {
	Id   int64  `gorm:"id" json:"id"`
	Code string `gorm:"code" json:"code"`
	Name string `gorm:"name" json:"name"`
}
