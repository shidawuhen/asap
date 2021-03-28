/**
@date: 2021/3/28
**/
package design

import (
	"fmt"
	"strconv"
)

type Water struct {
	weight int64
}

type Animal struct {
	name string
}

/**
 * @Description: 动物成长
 * @receiver animal
 * @param water
 */
func (animal *Animal) Grow(water *Water) {
	weight := strconv.FormatInt(water.weight, 10)
	fmt.Printf("长大需要喝%s吨水\n", weight)
}

type Wing struct {
}

/**
 * @Description: 鸟类，伪继承Animal，组合Wing，有颜色成员变量
 */
type Bird struct {
	Animal
	featherCol string
	wing       *Wing
}

/**
 * @Description:鸟下蛋
 * @receiver bird
 */
func (bird *Bird) LayEggs() {
	fmt.Println(bird.featherCol + "鸟下蛋了哈")
}

/**
 * @Description:飞接口
 */
type Fly interface {
	fly()
}

/**
 * @Description:大雁接口，伪继承鸟，并实现Fly接口
 */
type Goose struct {
	Bird
}

func (goose *Goose) fly() {
	fmt.Println("大雁可以飞")
}

/**
 * @Description:雁群聚合大雁
 */
type Geese struct {
	geese []*Goose
}

/**
 * @Description: 雁群的飞行方式
 * @receiver g
 */
func (g *Geese) flyV() {
	fmt.Printf("雁群V形飞行，共有%d只大雁\n", len(g.geese))
}

/**
 * @Description:鸭伪继承鸟
 */
type Duck struct {
	Bird
}

/**
 * @Description:说话接口
 */
type Speak interface {
	Speak()
}

/**
 * @Description:唐老鸭伪继承鸭子
 */
type DonaldDuck struct {
	Duck
}

/**
 * @Description:唐老鸭实现说话接口
 * @receiver donaldDuck
 */
func (donaldDuck *DonaldDuck) Speak() {
	fmt.Println("唐老鸭说人话了")
}

/**
 * @Description:企鹅伪继承鸟，并关联天气
 */
type Penguin struct {
	Bird
	climate *Climate
}

/**
 * @Description:企鹅重载了下蛋接口
 * @receiver p
 */
func (p *Penguin) LayEggs() {
	fmt.Printf("企鹅在%s的天气下下蛋了\n", p.climate.content)
}

/**
 * @Description:天气类，有天气内容成员变量
 */
type Climate struct {
	content string
}

func showUML() {
	water := &Water{weight: 10}
	//动物
	fmt.Println("-----------动物篇-依赖关系")
	animal := &Animal{}
	animal.Grow(water)
	//鸟
	fmt.Println("-----------鸟篇-伪继承/组合关系")
	bird := &Bird{
		featherCol: "五彩斑斓的",
		wing:       &Wing{},
	}
	bird.Grow(water)
	bird.LayEggs()
	//大雁
	fmt.Println("-----------大雁篇-实现关系")
	goose := &Goose{}
	goose.featherCol = "黑黑的"
	goose.Grow(water)
	goose.LayEggs()
	goose.fly()
	//鸭
	fmt.Println("-----------鸭篇-伪继承关系")
	duck := &Duck{}
	duck.featherCol = "黄色的"
	duck.Grow(water)
	duck.LayEggs()
	//企鹅
	fmt.Println("-----------企鹅篇-伪继承/关联关系")
	penguin := &Penguin{}
	penguin.featherCol = "白色的"
	penguin.Grow(water)
	climate := &Climate{content: "寒冷的"}
	penguin.climate = climate
	penguin.LayEggs()
	//雁群
	fmt.Println("-----------雁群篇-聚合关系")
	g := &Geese{}
	g.geese = append(g.geese, goose)
	g.flyV()
	//唐老鸭
	fmt.Println("-----------唐老鸭篇-实现关系")
	donaldDuck := &DonaldDuck{}
	donaldDuck.featherCol = "红色的"
	donaldDuck.Grow(water)
	donaldDuck.LayEggs()
	donaldDuck.Speak()
}
