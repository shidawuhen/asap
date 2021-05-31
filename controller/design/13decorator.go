/**
@date: 2021/5/28
**/
package design

import "fmt"

/**
 * @Description: 飞行器接口，有fly函数
 */
type Aircraft interface {
	fly()
	landing()
}

/**
 * @Description: 直升机类，拥有正常飞行、降落功能
 */
type Helicopter struct {
}

func (h *Helicopter) fly() {
	fmt.Println("我是普通直升机")
}

func (h *Helicopter) landing() {
	fmt.Println("我有降落功能")
}

/**
 * @Description: 武装直升机
 */
type WeaponAircraft struct {
	Aircraft
}

/**
 * @Description: 给直升机增加武装功能
 * @receiver a
 */
func (a *WeaponAircraft) fly() {
	a.Aircraft.fly()
	fmt.Println("增加武装功能")
}

/**
 * @Description: 救援直升机
 */
type RescueAircraft struct {
	Aircraft
}

/**
 * @Description: 给直升机增加救援功能
 * @receiver r
 */
func (r *RescueAircraft) fly() {
	r.Aircraft.fly()
	fmt.Println("增加救援功能")
}

func decoratorMain() {
	//普通直升机
	fmt.Println("------------普通直升机")
	helicopter := &Helicopter{}
	helicopter.fly()
	helicopter.landing()

	//武装直升机
	fmt.Println("------------武装直升机")
	weaponAircraft := &WeaponAircraft{
		Aircraft: helicopter,
	}
	weaponAircraft.fly()

	//救援直升机
	fmt.Println("------------救援直升机")
	rescueAircraft := &RescueAircraft{
		Aircraft: helicopter,
	}
	rescueAircraft.fly()

	//武装救援直升机
	fmt.Println("------------武装救援直升机")
	weaponRescueAircraft := &RescueAircraft{
		Aircraft: weaponAircraft,
	}
	weaponRescueAircraft.fly()
}
