/**
@date: 2021/5/1
**/
package design

import (
	"errors"
	"fmt"
)

/**
 * @Description: Product内的参数
 */
type ResourceParams struct {
	name     string
	maxTotal int64
	maxIdle  int64
	minIdle  int64
}

/**
 * @Description: Product接口
 */
type ResourceProduct interface {
	show()
}

/**
 * @Description: 实际Product，有show函数
 */
type RedisResourceProduct struct {
	resourceParams ResourceParams
}

/**
 * @Description: show成员函数，用于显示product的参数内容
 * @receiver p
 */
func (p *RedisResourceProduct) show() {
	fmt.Printf("Product的数据为 %+v ", p.resourceParams)
}

/**
 * @Description: 资源类创建接口
 */
type ResourceBuilder interface {
	setName(name string) ResourceBuilder
	setMaxTotal(maxTotal int64) ResourceBuilder
	setMaxIdle(maxIdle int64) ResourceBuilder
	setMinIdle(minIdle int64) ResourceBuilder
	getError() error
	build() (p ResourceProduct)
}

/**
 * @Description: 实际建造者
 */
type RedisResourceBuilder struct {
	resourceParams ResourceParams
	err            error
}

/**
 * @Description: 获取错误信息
 * @receiver r
 * @return error
 */
func (r *RedisResourceBuilder) getError() error {
	return r.err
}

/**
 * @Description: 设置名称
 * @receiver r
 * @param name
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setName(name string) ResourceBuilder {
	if name == "" {
		r.err = errors.New("name为空")
		return r
	}
	r.resourceParams.name = name
	fmt.Println("RedisResourceBuilder setName ", name)
	return r
}

/**
 * @Description: 设置maxTotal值，值不能小于0
 * @receiver r
 * @param maxTotal
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMaxTotal(maxTotal int64) ResourceBuilder {
	if maxTotal <= 0 {
		r.err = errors.New("maxTotal小于0")
		return r
	}
	r.resourceParams.maxTotal = maxTotal
	fmt.Println("RedisResourceBuilder setMaxTotal ", maxTotal)
	return r
}

/**
 * @Description: 设置maxIdle值，值不能小于0
 * @receiver r
 * @param maxIdle
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMaxIdle(maxIdle int64) ResourceBuilder {
	if maxIdle <= 0 {
		r.err = errors.New("maxIdle小于0")
		return r
	}
	r.resourceParams.maxIdle = maxIdle
	fmt.Println("RedisResourceBuilder setMaxIdle ", maxIdle)
	return r
}

/**
 * @Description: 设置minIdle值，值不能小于0
 * @receiver r
 * @param minIdle
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMinIdle(minIdle int64) ResourceBuilder {
	if minIdle <= 0 {
		r.err = errors.New("minIdle小于0")
		return r
	}
	r.resourceParams.minIdle = minIdle
	fmt.Println("RedisResourceBuilder setMinIdle ", minIdle)
	return r
}

/**
 * @Description: 构建product
	1. 做参数校验
	2. 根据参数生成product
 * @receiver r
 * @return p
*/
func (r *RedisResourceBuilder) build() (p ResourceProduct) {
	// 校验逻辑放到这里来做，包括必填项校验、依赖关系校验、约束条件校验等
	if r.resourceParams.name == "" {
		r.err = errors.New("name为空")
		return
	}

	if !((r.resourceParams.maxIdle == 0 && r.resourceParams.minIdle == 0 && r.resourceParams.maxTotal == 0) ||
		(r.resourceParams.maxIdle != 0 && r.resourceParams.minIdle != 0 && r.resourceParams.maxTotal != 0)) {
		r.err = errors.New("数据需要保持一致")
		return
	}

	if r.resourceParams.maxIdle > r.resourceParams.maxTotal {
		r.err = errors.New("maxIdle > maxTotal")
		return
	}
	if r.resourceParams.minIdle > r.resourceParams.maxTotal || r.resourceParams.minIdle > r.resourceParams.maxIdle {
		r.err = errors.New("minIdle > maxTotal|maxIdle")
		return
	}
	fmt.Println("RedisResourceBuilder build")
	product := &RedisResourceProduct{
		resourceParams: r.resourceParams,
	}
	return product
}

/**
 * @Description: 指挥者
 */
type Director struct {
}

/**
 * @Description: 指挥者控制建造过程
 * @receiver d
 * @param builder
 * @return *ResourceProduct
 */
func (d *Director) construct(builder ResourceBuilder) ResourceProduct {
	resourceProduct := builder.setName("redis").
		setMinIdle(10).
		setMaxIdle(10).
		setMaxTotal(20).
		build()

	err := builder.getError()
	if err != nil {
		fmt.Println("构建失败，原因为" + err.Error())
		return nil
	}
	return resourceProduct
}

func builderMain() {
	builder := &RedisResourceBuilder{}

	director := &Director{}
	product := director.construct(builder)

	if product == nil {
		return
	}

	product.show()
}
