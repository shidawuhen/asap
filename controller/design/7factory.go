/**
@author: pangzhiqiang
@date: 2021/4/25
**/
package design

import "fmt"

/**
 * @Description: 解析接口，有解析函数用来解析文件
 */
type ConfigParse interface {
	Parse(path string) string
}

/**
 * @Description: Json解析类，实现解析接口
 */
type JsonConfigParse struct {
}

/**
 * @Description: 用于解析Json文件
 * @receiver json
 * @param path
 * @return string
 */
func (json *JsonConfigParse) Parse(path string) string {
	return "解析json配置文件，路径为:" + path
}

/**
 * @Description: Yaml解析类，实现解析接口
 */
type YamlConfigParse struct {
}

/**
 * @Description: 用于解析Yaml文件
 * @receiver yaml
 * @param path
 * @return string
 */
func (yaml *YamlConfigParse) Parse(path string) string {
	return "解析yaml配置文件，路径为:" + path
}

/**
 * @Description: 简单工厂
 */
type SimpleParseFactory struct {
}

func (simple *SimpleParseFactory) create(ext string) ConfigParse {
	switch ext {
	case "json":
		return &JsonConfigParse{}
	case "yaml":
		return &YamlConfigParse{}
	}
	return nil
}

/**
 * @Description: 工厂方法
 */
type NormalParseFactory interface {
	createParse() ConfigParse
}

/**
 * @Description: Json工厂方法
 */
type JsonNormalParseFactory struct {
}

/**
 * @Description: 该方法用于创建Json解析器
 * @receiver jsonFactory
 * @param ext
 * @return ConfigParse
 */
func (jsonFactory *JsonNormalParseFactory) createParse() ConfigParse {
	//此处假装有各种组装
	return &JsonConfigParse{}
}

/**
 * @Description: Yaml解析工厂
 */
type YamlNormalParseFactory struct {
}

/**
 * @Description: 该方法用于创建Yaml解析器
 * @receiver yamlFactory
 * @return ConfigParse
 */
func (yamlFactory *YamlNormalParseFactory) createParse() ConfigParse {
	//此处假装有各种组装
	return &YamlConfigParse{}
}

/**
 * @Description: 根据后缀创建工厂
 * @param ext
 * @return NormalParseFactory
 */
func createFactory(ext string) NormalParseFactory {
	switch ext {
	case "json":
		return &JsonNormalParseFactory{}
	case "yaml":
		return &YamlNormalParseFactory{}
	}
	return nil
}

func factoryShow() {
	//简单工厂使用代码
	fmt.Println("------------简单工厂")
	simpleParseFactory := &SimpleParseFactory{}
	parse := simpleParseFactory.create("json")
	if parse != nil {
		data := parse.Parse("conf/config.json")
		fmt.Println(data)
	}
	//工厂使用代码
	fmt.Println("------------工厂方法")
	factory := createFactory("yaml")
	parse = factory.createParse()
	if parse != nil {
		data := parse.Parse("conf/config.yaml")
		fmt.Println(data)
	}
}
