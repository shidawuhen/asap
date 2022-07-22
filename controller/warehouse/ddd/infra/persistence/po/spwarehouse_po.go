/**
@author: Jason Pang
@desc:
@date: 2022/7/22
**/
package po

import "time"

type SpWareHouse struct {
	Id            int64     `gorm:"id" json:"id"`
	Code          string    `gorm:"code" json:"code"`
	Name          string    `gorm:"name" json:"name"`
	SpId          int64     `gorm:"column:scsp_id" json:"scsp_id"`
	WhType        int64     `gorm:"wh_type" json:"wh_type"`
	ResourceType  string    `gorm:"resource_type" json:"resource_type"`
	Addr          string    `gorm:"addr" json:"addr"`
	ContactName   string    `gorm:"contact_name" json:"contact_name"`
	ContactNumber string    `grom:"contact_number" json:"contact_number"`
	IsActive      int64     `gorm:"is_active" json:"is_active"`
	Updater       string    `gorm:"column:update_by" json:"update_by"`
	UpdateTime    time.Time `gorm:"update_time" json:"update_time"`
	CustomsId     int64     `gorm:"customs_id" json:"customs_id"`
	CustomsCode   string    `gorm:"customs_code" json:"customs_code"`
	ZipCode       string    `gorm:"zip_code" json:"zip_code"`
	Imgs          string    `gorm:"imgs" json:"imgs"`
	Mode          string    `gorm:"mode" json:"mode"`
	AddressName   string    `gorm:"address_name" json:"address_name"`
	AddressCode   string    `gorm:"address_code" json:"address_code"`
}
