/**
@date: 2021/5/14
**/
package design

import (
	"fmt"
)

/**
 * @Description: 支付接口，只包含发起支付功能
 */
type PaymentService interface {
	pay(order string) string
}

/**
 * @Description: 微信支付类
 */
type WXPay struct {
}

/**
 * @Description: 微信支付类，从微信获取支付token
 * @receiver w
 * @param order
 * @return string
 */
func (w *WXPay) pay(order string) string {
	return "从微信获取支付token"
}

/**
 * @Description: 阿里支付类
 */
type AliPay struct {
}

/**
 * @Description: 阿里支付类，从阿里获取支付token
 * @receiver a
 * @param order
 * @return string
 */
func (a *AliPay) pay(order string) string {
	return "从阿里获取支付token"
}

/**
 * @Description: 支付代理类
 */
type PaymentProxy struct {
	realPay PaymentService
}

/**
 * @Description: 做校验签名、初始化订单数据、参数检查、记录日志、组装这种通用性操作，调用真正支付类获取token
 * @receiver p
 * @param order
 * @return string
 */
func (p *PaymentProxy) pay(order string) string {
	fmt.Println("处理" + order)
	fmt.Println("1校验签名")
	fmt.Println("2格式化订单数据")
	fmt.Println("3参数检查")
	fmt.Println("4记录请求日志")
	token := p.realPay.pay(order)
	return "http://组装" + token + "然后跳转到第三方支付"
}
func proxyMain() {
	proxy := &PaymentProxy{
		realPay: &AliPay{},
	}
	url := proxy.pay("阿里订单")
	fmt.Println(url)
}
