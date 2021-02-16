package design

import "fmt"

//继承
type Base struct {
	Name string
}

func (b *Base) SetName(name string) {
	b.Name = name
}

func (b *Base) GetName() string {
	return b.Name
}

// 组合，实现继承
type Child struct {
	base Base // 这里保存的是Base类型
}

// 重写GetName方法
func (c *Child) GetName() string {
	c.base.SetName("modify...")
	return c.base.GetName()
}

// 实现继承，但需要外部提供一个Base的实例
type Child2 struct {
	base *Base // 这里是指针
}

//
type Child3 struct {
	Base
}

type Child4 struct {
	*Base
}

//多态
type Money interface {
	show() string
}

type OldMoney struct {
}

func (oldMoney *OldMoney) show() string {
	return "I am old money"
}

type NewMoney struct {
}

func (newMoney *NewMoney) show() string {
	return "I am new money"
}

func PrintMoney(l []Money) {
	for _, item := range l {
		fmt.Println(item.show())
	}
}

func mainprammar() {
	//继承
	c := new(Child)
	c.base.SetName("world")
	fmt.Println(c.GetName())

	c2 := new(Child2)
	c2.base = new(Base) // 因为Child2里面的Base是指针类型，所以必须提供一个Base的实例
	c2.base.SetName("ccc")
	fmt.Println(c2.base.GetName())

	c3 := new(Child3)
	c3.SetName("1111")
	fmt.Println(c3.GetName())

	c4 := new(Child4)
	c4.Base = new(Base)
	c4.SetName("2222")
	fmt.Println(c4.GetName())
	//多态
	moneyList := []Money{new(OldMoney), new(NewMoney), new(OldMoney)}
	PrintMoney(moneyList)
}