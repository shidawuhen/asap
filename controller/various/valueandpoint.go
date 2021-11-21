/**
@author: Jason Pang
@desc: Go值和指针
@date: 2021/11/21
**/
package various

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type ImageItem struct {
	Key      string `thrift:"key,1" json:"key"`
	Name     string `thrift:"name,2" json:"name"`
	FileType string `thrift:"fileType,3" json:"fileType"`
}
type Test struct {
	Hello string
	World string
}

//错误示例
func ErrorShow() {
	fmt.Println("------------------错误示例")
	var a, b string = "a", "b"
	ai := ImageItem{
		Key:      a,
		Name:     a,
		FileType: a,
	}
	bi := ImageItem{
		Key:      b,
		Name:     b,
		FileType: b,
	}
	//news
	newS := make(map[int]*string)
	newL := []ImageItem{ai, bi}
	for k, item := range newL {
		fmt.Println("------------------第", k+1, "次循环")
		fmt.Printf("item的地址为 %p\n", &item)
		fmt.Println("item的值为", item)
		fmt.Printf("新key的地址为%p,老key的地址为%p\n", &item.Key, &newL[k].Key)
		newS[k] = &item.Key
		fmt.Printf("newS[k]的地址为: %p\n", newS[k])
		fmt.Printf("newS[k]的值为: %s\n", *newS[k])
	}
	f, _ := json.Marshal(newS)
	fmt.Println(string(f))
}

func ValueAndPoint() {
	fmt.Println("------------------简单变量")
	var hello string = "hello"
	fmt.Println("值为", hello)
	fmt.Printf("字符串字节长度为%d\n", unsafe.Sizeof(hello))
	fmt.Println("地址为", &hello)
	fmt.Printf("地址为%p\n", &hello)
	//结构体
	t := Test{
		Hello: "hello",
		World: "world",
	}
	fmt.Println("------------------结构体")
	fmt.Println("值为", t)
	fmt.Printf("结构体字节长度为%d\n", unsafe.Sizeof(t))
	fmt.Println("地址为", &t)
	fmt.Printf("地址为%p\n", &t)
	fmt.Printf("地址为%p\n", &t.Hello)
	fmt.Printf("地址为%p\n", &t.World)
	//地址存放
	fmt.Println("------------------地址存放")
	var save *Test = &t
	fmt.Printf("值为%p\n", &*save)
	fmt.Printf("指针字节长度为%d\n", unsafe.Sizeof(save))
	fmt.Printf("地址为%p\n", &save)
}

func NewCase() {
	fmt.Println("------------------正确")
	var a, b string = "a", "b"
	ai := ImageItem{
		Key:      a,
		Name:     a,
		FileType: a,
	}
	bi := ImageItem{
		Key:      b,
		Name:     b,
		FileType: b,
	}
	l := make(map[int]*ImageItem)
	l[0] = &ai
	l[1] = &bi
	s := make([]*ImageItem, 2)
	//指针赋值，没影响
	for k, item := range l {
		fmt.Println("------------------第", k+1, "次循环")
		fmt.Printf("原数据的指针地址为: %p\n", l[k])
		fmt.Printf("原数据的数据为: %v\n", l[k])

		fmt.Println("item的地址为", &item)
		fmt.Printf("item的值为: %p\n", item)

		s[k] = item

		fmt.Println("s[k]的地址为", &s[k])
		fmt.Printf("s[k]的值为: %p\n", s[k])
	}

	f, _ := json.Marshal(s)
	fmt.Println(string(f))
}

func NormalErrorCase() {
	fmt.Println("------------------错误")
	var a, b string = "a", "b"
	ai := ImageItem{
		Key:      a,
		Name:     a,
		FileType: a,
	}
	bi := ImageItem{
		Key:      b,
		Name:     b,
		FileType: b,
	}
	l := make(map[int]ImageItem)
	l[0] = ai
	l[1] = bi
	s := make([]*ImageItem, 2)
	//指针赋值，没影响
	for k, item := range l {
		s[k] = &item
	}

	f, _ := json.Marshal(s)
	fmt.Println(string(f))
}

func valueAndPointMain() {
	//ErrorShow()
	//ValueAndPoint()
	//NewCase()
	NormalErrorCase()
}
