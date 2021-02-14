package design

import "fmt"

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
	moneyList := []Money{new(OldMoney), new(NewMoney), new(OldMoney)}
	PrintMoney(moneyList)
}