/**
@author: Jason Pang
@desc:
@date: 2022/7/4
**/
package idl

type ShopWareHouseInfo struct {
	Id            int64  `json:"id"`
	WarehouseId   int64  `json:"warehouse_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	SpWareHouseId int64  `json:"scsp_warehouse_id"`
	SpCode        string `json:"sp_code"`
	SpName        string `json:"sp_name"`
}
