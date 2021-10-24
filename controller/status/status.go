/**
@author: Jason Pang
@desc:
@date: 2021/10/24
**/
package status

import (
	"fmt"
	"reflect"
)

// 时序图
var StatusTimingGraph = map[string][]string{
	"A":  {"B1", "B2"},
	"B1": {"C1", "C2"},
	"B2": {"B1"},
	"C1": {"D"},
	"C2": {"C1"},
	"D":  {"E"},
}

// 核心节点
var CoreStatus = []string{
	"A",
	"B1",
	"C1",
	"E",
}
var StatusJumpGraph = InitJumpGraph(StatusTimingGraph, CoreStatus)

func InitJumpGraph(statusMap map[string][]string, coreStatus []string) map[string][]string {
	retMap := make(map[string][]string, 0)
	for status, statusList := range statusMap {
		retList := make([]string, 0)
		for _, tStatus := range statusList {
			retList = append(retList, tStatus)
			if InSlice(coreStatus, tStatus) {
				continue
			}
			tList := recursionGraph(tStatus, statusMap, coreStatus)
			for _, tStatus := range tList {
				if !InSlice(retList, tStatus) {
					retList = append(retList, tStatus)
				}
			}
		}

		retMap[status] = retList
	}
	return retMap
}
func recursionGraph(status string, statusMap map[string][]string, coreStatus []string) []string {
	retList := make([]string, 0)
	if statusList, ok := statusMap[status]; ok {
		for _, tStatus := range statusList {
			retList = append(retList, tStatus)
			if InSlice(coreStatus, tStatus) {
				continue
			}
			retList = append(retList, recursionGraph(tStatus, statusMap, coreStatus)...)
		}
	}
	return retList
}
func InSlice(a, b interface{}) bool {
	exist, _ := InSliceWithError(a, b)
	return exist
}
func InSliceWithError(a, b interface{}) (exist bool, err error) {

	va := reflect.ValueOf(a)

	if va.Kind() != reflect.Slice {
		err = fmt.Errorf("parameter a must be a slice")
		return
	}

	if reflect.TypeOf(a).String()[2:] != reflect.TypeOf(b).String() {
		err = fmt.Errorf("type of parameter b not match with parameter a")
		return
	}

	for i := 0; i < va.Len(); i++ {
		if va.Index(i).Interface() == b {
			exist = true
			return
		}
	}

	return
}

func statusMain() {
	originStatus := "A"
	targetStatus := "B1"
	statusList, ok := StatusJumpGraph[originStatus]
	fmt.Println(StatusJumpGraph)
	if !ok {
		fmt.Println("状态有误")
		return
	}

	if !InSlice(statusList, targetStatus) {
		fmt.Println("状态不合规，无法流转")
		return
	}
}
