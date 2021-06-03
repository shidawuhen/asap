/**
@author: Jason Pang
@date: 2021/6/3
**/
package design

import (
	"fmt"
	"time"
)

/**
 * @Author: Jason Pang
 * @Description: 对账单数据
 */
type StatementItem struct {
	OrderId       string //系统单号
	TransactionId string //第三方交易号
	Amount        int64  //支付金额，单位：分
	PaymentTime   int64  //订单支付时间
}

/**
 * @Author: Jason Pang
 * @Description: 从第三方获取对账数据
 */
type StatementData interface {
	GetStatementData(startTime int64, endTime int64) []*StatementItem
}

/**
 * @Author: Jason Pang
 * @Description: WX支付
 */
type WXStatementData struct {
}

func (w *WXStatementData) GetStatementData(startTime int64, endTime int64) []*StatementItem {
	fmt.Println("从WX获取到的对账数据，支付时间需要格式化为时间戳")
	return []*StatementItem{
		{
			OrderId:       "WX订单222",
			TransactionId: "WX支付单号",
			Amount:        999,
			PaymentTime:   time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local).Unix(),
		},
	}
}

/**
 * @Author: Jason Pang
 * @Description: ZFB支付
 */
type ZFBStatementData struct {
}

func (z *ZFBStatementData) GetStatementData(startTime int64, endTime int64) []*StatementItem {
	fmt.Println("从ZFB获取到的对账数据，金额需要从元转化为分")
	return []*StatementItem{
		{
			OrderId:       "ZFB订单111",
			TransactionId: "ZFB支付单号",
			Amount:        99.9 * 100,
			PaymentTime:   1389058332,
		},
	}
}

/**
 * @Author: Jason Pang
 * @Description: 对账函数
 * @param list  从第三方获取的对账单
 * @return bool
 */
func DoStatement(list []*StatementItem) bool {
	fmt.Println("开始对账")
	fmt.Println("从自身系统中获取指定时间内的支付单")
	for _, item := range list {
		fmt.Println(item.OrderId + " 与系统支付单进行对账")
	}
	fmt.Println("对账完成")
	return true
}

func adapterMain() {
	wx := &WXStatementData{}
	zfb := &ZFBStatementData{}
	stattementData := []StatementData{
		wx,
		zfb,
	}
	for _, s := range stattementData {
		DoStatement(s.GetStatementData(1389058332, 1389098332))
	}
}
