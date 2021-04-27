/**
@date: 2021/4/27
**/
package design

import "fmt"

/////////////////////////////////////////////////// Product部分
/**
 * @Description: 学生接口，定义了插入和更新功能
 */
type Student interface {
	insert() bool
	update() bool
}

/**
 * @Description: Access操作Student类
 */
type AccessStudent struct {
}

/**
 * @Description: 使用Access向Student表中插入数据
 * @receiver a
 * @return bool
 */
func (a *AccessStudent) insert() bool {
	fmt.Println("AccessStudent insert")
	return true
}

/**
 * @Description: 使用Access向Student表中更新数据
 * @receiver a
 * @return bool
 */
func (a *AccessStudent) update() bool {
	fmt.Println("AccessStudent update")
	return true
}

/**
 * @Description: MySQL操作Student类
 */
type MySQLStudent struct {
}

/**
 * @Description: 使用MySQL向Student表中插入数据
 * @receiver a
 * @return bool
 */
func (m *MySQLStudent) insert() bool {
	fmt.Println("MySQLStudent insert")
	return true
}

/**
 * @Description: 使用MySQL向Student表中更新数据
 * @receiver a
 * @return bool
 */
func (m *MySQLStudent) update() bool {
	fmt.Println("MySQLStudent update")
	return true
}

/**
 * @Description: 成绩接口，定义了插入和列表功能
 */
type Score interface {
	insert() bool
	list() []int64
}

/**
 * @Description: 使用Access操作Score类
 */
type AccessScore struct {
}

/**
 * @Description: 使用Access向Score表中插入数据
 * @receiver a
 * @return bool
 */
func (a *AccessScore) insert() bool {
	fmt.Println("AccessScore insert")
	return true
}

/**
 * @Description: 使用Access从Score表中获取成绩列表
 * @receiver a
 * @return []int64
 */
func (a *AccessScore) list() []int64 {
	fmt.Println("AccessScore list")
	return []int64{1, 2}
}

/**
 * @Description: 使用MySQL操作Score类
 */
type MySQLScore struct {
}

/**
 * @Description: 使用MySQL向Score表中插入数据
 * @receiver a
 * @return bool
 */
func (m *MySQLScore) insert() bool {
	fmt.Println("MySQLScore insert")
	return true
}

/**
 * @Description: 使用MySQL从Score表中获取成绩列表
 * @receiver a
 * @return []int64
 */
func (m *MySQLScore) list() []int64 {
	fmt.Println("MySQLScore list")
	return []int64{1, 2}
}

/////////////////////////////////////////////////// Factory部分
/**
 * @Description: 抽象工厂接口，代表高维度工厂，高维度工厂能够生成低维度对象
 */
type Factory interface {
	createStudent() Student
	createScore() Score
}

/**
 * @Description: 高维度Access工厂
 */
type AccessFactory struct {
}

/**
 * @Description: 高维度Access工厂，创建Access的Student对象
 * @receiver a
 * @return Student
 */
func (a *AccessFactory) createStudent() Student {
	return &AccessStudent{}
}

/**
 * @Description: 高维度Access工厂，创建Access的Score对象
 * @receiver a
 * @return Score
 */
func (a *AccessFactory) createScore() Score {
	return &AccessScore{}
}

/**
 * @Description: 高维度MySQL工厂
 */
type MySQLFactory struct {
}

/**
 * @Description: 高维度MySQL工厂，创建MySQL的Student对象
 * @receiver a
 * @return Student
 */
func (m *MySQLFactory) createStudent() Student {
	return &MySQLStudent{}
}

/**
 * @Description: 高维度MySQL工厂，创建MySQL的Score对象
 * @receiver a
 * @return Score
 */
func (m *MySQLFactory) createScore() Score {
	return &MySQLScore{}
}

/////////////////////////////////////////////////// 获得高维度工厂
func getFactory(storeType string) Factory {
	switch storeType {
	case "MySQL":
		return &MySQLFactory{}
	case "Access":
		return &AccessFactory{}
	}
	return nil
}

func abstractFactory() {
	//抽象工厂使用代码
	fmt.Println("------------抽象工厂")
	factory := getFactory("MySQL")
	if factory == nil {
		fmt.Println("不支持该存储方式")
		return
	}

	student := factory.createStudent()
	score := factory.createScore()

	student.insert()
	student.update()
	score.insert()
	score.list()
}
