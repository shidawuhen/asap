/**
@author: Jason Pang
@desc:
@date: 2021/7/7
**/
package design

import "fmt"

const (
	Common = "COMMON"
	Win    = "WIN"
)

/**
 * @Author: Jason Pang
 * @Description: 根据hscode获取商品类型
 * @param hscode
 * @return string
 */
func getProductType(hscode string) string {
	if hscode == "11" {
		return Common
	} else {
		return Win
	}
}

/**
 * @Author: Jason Pang
 * @Description: 税费计算函数，金额都为分
 * @param price
 * @param qty
 * @return taxPrice
 */
type TaxComputeFunc func(price int64, qty int64) (taxPrice int64)

/**
 * @Author: Jason Pang
 * @Description: 税费计算策略存储处
	0为不含税 1为含税
*/
var TaxComputeFuncMap = map[int]map[string]TaxComputeFunc{
	0: map[string]TaxComputeFunc{
		Common: common,
		Win:    win,
	},
	1: map[string]TaxComputeFunc{
		Common: common,
		Win:    win,
	},
}

/**
 * @Author: Jason Pang
 * @Description: 计算普通商品税费
 * @param price
 * @param qty
 * @return taxPrice
 */
func common(price int64, qty int64) (taxPrice int64) {
	radio := 0.1
	fmt.Println("计算普通商品税费")
	return int64(float64(price*qty) * radio)
}

/**
 * @Author: Jason Pang
 * @Description: 计算酒类税费
 * @param price
 * @param qty
 * @return taxPrice
 */
func win(price int64, qty int64) (taxPrice int64) {
	radio := 0.2
	fmt.Println("计算普酒类税费")
	return int64(float64(price*qty) * radio)
}

/**
 * @Author: Jason Pang
 * @Description: 计算税费
 * @param withTax
 * @param productType
 * @param price
 * @param qty
 */
func ComputeTaxPrice(withTax int, productType string, price int64, qty int64) {
	if taxFunc, ok := TaxComputeFuncMap[withTax][productType]; ok {
		taxPrice := taxFunc(price, qty)
		fmt.Println("税费为", taxPrice)
	} else {
		fmt.Println("输入有误，无法计算")
	}
}

func strategyMain() {
	//获取商品是否含税、商品价格、商品数量、商品类型
	withTax := 0
	var price, qty int64 = 10000, 3
	productType := getProductType("11")
	//计算税费
	ComputeTaxPrice(withTax, productType, price, qty)
}
