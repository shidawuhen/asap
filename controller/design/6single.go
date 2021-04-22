/**
@date: 2021/4/22
单例模式
**/
package design

import (
	"fmt"
	"sync"
)

type singleTon struct {
}

func (s *singleTon) Show() {
	fmt.Println("hello world")
}

var (
	once   sync.Once
	single *singleTon
)

func GetSingleInstance() *singleTon {
	once.Do(func() {
		single = &singleTon{}
	})
	return single
}
