/**
@author: Jason Pang
@desc:
@date: 2021/8/6
**/
package design

import (
	"fmt"
	"math/rand"
)

/**
 * @Author: Jason Pang
 * @Description: 二维数组，当做图集合，写个真正的图集合太麻烦了
 */
type TwoDimensionalArray struct {
	row   int64
	col   int64
	array [][]int64
}

/**
 * @Author: Jason Pang
 * @Description: 设置尺寸
 * @receiver t
 * @param row
 * @param col
 */
func (t *TwoDimensionalArray) SetSize(row int64, col int64) {
	t.row = row
	t.col = col
	t.array = make([][]int64, row)
	var i int64 = 0
	for i = 0; i < row; i++ {
		t.array[i] = make([]int64, col)
	}
}

/**
 * @Author: Jason Pang
 * @Description: 初始化数据
 * @receiver t
 */
func (t *TwoDimensionalArray) InitDefault() {
	if t.row <= 0 || t.col <= 0 {
		return
	}
	var i, j int64 = 0, 0
	for i = 0; i < t.row; i++ {
		for j = 0; j < t.col; j++ {
			t.array[i][j] = rand.Int63n(200)
		}
	}
}

/**
 * @Author: Jason Pang
 * @Description: 格式化输出。不能为指针。
 * @receiver t
 * @return string
 */
func (t TwoDimensionalArray) String() string {
	s := ""
	var i int64 = 0
	for i = 0; i < t.row; i++ {
		s += fmt.Sprintf("%v \n", t.array[i])
	}
	return s
}

/**
 * @Author: Jason Pang
 * @Description: 迭代器接口
 */
type Iterator interface {
	First()
	IsDone() bool
	CurrentItem()
	Next()
}

type BFSPos struct {
	x, y int64
}

/**
 * @Author: Jason Pang
 * @Description: 广度优先遍历
 */
type BFSIterator struct {
	data     *TwoDimensionalArray
	used     [][]bool
	queue    []BFSPos
	index    int64 //queue遍历的位置
	bfsIndex int64 //记录做广度优先遍历的位置
}

/**
 * @Author: Jason Pang
 * @Description: 赋值
 * @receiver d
 * @param data
 */
func (d *BFSIterator) Create(data *TwoDimensionalArray) {
	d.data = data
	var i int64 = 0
	d.used = make([][]bool, data.row)
	for i = 0; i < data.row; i++ {
		d.used[i] = make([]bool, data.col)
	}
	d.used[0][0] = true
	d.queue = make([]BFSPos, 1)
	d.queue[0] = BFSPos{0, 0}
	d.index = 0
	d.bfsIndex = 0
}

/**
 * @Author: Jason Pang
 * @Description: 初始数据
 * @receiver d
 */
func (d *BFSIterator) First() {
	fmt.Println(d.data.array[0][0])
}

/**
 * @Author: Jason Pang
 * @Description: 是否遍历结束
 * @receiver d
 * @return bool
 */
func (d *BFSIterator) IsDone() bool {
	if d.index == d.data.col*d.data.row {
		return true
	}
	return false
}

/**
 * @Author: Jason Pang
 * @Description: 当前数值
 * @receiver d
 */
func (d *BFSIterator) CurrentItem() {
	pos := d.queue[d.index]
	fmt.Println(d.index, ":", d.data.array[pos.x][pos.y])
}

/**
 * @Author: Jason Pang
 * @Description: 移动
 * @receiver d
 */
func (d *BFSIterator) Next() {
	if d.index >= d.data.row*d.data.col {
		fmt.Println("已到最后")
		return
	}
	//说明已经没有了，需要再加几个
	if d.index >= int64(len(d.queue))-1 {
		for d.bfsIndex < int64(len(d.queue)) && d.index < int64(len(d.queue)) {
			curI, curJ := d.queue[d.bfsIndex].x, d.queue[d.bfsIndex].y
			if curJ+1 < d.data.col && d.used[curI][curJ+1] == false {
				d.queue = append(d.queue, BFSPos{curI, curJ + 1})
				d.used[curI][curJ+1] = true
			}
			if curI+1 < d.data.row && curJ+1 < d.data.col && d.used[curI+1][curJ+1] == false {
				d.queue = append(d.queue, BFSPos{curI + 1, curJ + 1})
				d.used[curI+1][curJ+1] = true
			}
			if curI+1 < d.data.row && d.used[curI+1][curJ] == false {
				d.queue = append(d.queue, BFSPos{curI + 1, curJ})
				d.used[curI+1][curJ] = true
			}
			d.bfsIndex++
		}

	}
	d.index++
}

func main() {
	t := TwoDimensionalArray{}
	t.SetSize(3, 3)
	t.InitDefault()
	fmt.Printf("%s", t)

	iterator := BFSIterator{}
	iterator.Create(&t)
	for iterator.IsDone() != true {
		iterator.CurrentItem()
		iterator.Next()
	}
}
