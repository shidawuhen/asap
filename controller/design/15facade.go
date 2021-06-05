/**
@author: Jason Pang
@date: 2021/6/5
**/
package design

import "fmt"

type ProductSystem struct {
}

func (p *ProductSystem) GetProductInfo() {
	fmt.Println("获取到商品信息")
}

type StockSystem struct {
}

func (s *StockSystem) GetStockInfo() {
	fmt.Println("获取到库存信息")
}

type PromotionSystem struct {
}

func (p *PromotionSystem) GetPromotionInfo() {
	fmt.Println("获取营销信息")
}

func ProductDetail() {
	product := &ProductSystem{}
	stock := &StockSystem{}
	promotion := &PromotionSystem{}
	product.GetProductInfo()
	stock.GetStockInfo()
	promotion.GetPromotionInfo()
	fmt.Println("整理完成商品详情页所有数据")
}
func facadeMain() {
	ProductDetail()
}
