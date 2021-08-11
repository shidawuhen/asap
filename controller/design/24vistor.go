/**
@author: Jason Pang
@desc:
@date: 2021/8/11
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description: 读文件接口，用于获取到文件
 */
type ReadFile interface {
	Read(fileName string)
	Accept(v VistorReadFile)
}

/**
 * @Author: Jason Pang
 * @Description: 读pdf文件类
 */
type ReadPdfFile struct {
}

/**
 * @Author: Jason Pang
 * @Description: 读取文件
 * @receiver p
 * @param fileName
 */
func (p *ReadPdfFile) Read(fileName string) {
	fmt.Println("读取pdf文件" + fileName)
}

/**
 * @Author: Jason Pang
 * @Description: 接受访问者类
 * @receiver p
 * @param v
 */
func (p *ReadPdfFile) Accept(v VistorReadFile) {
	v.VistorPdfFile(p)
}

/**
 * @Author: Jason Pang
 * @Description: 读取txt文件类
 */
type ReadTxtFile struct {
}

/**
 * @Author: Jason Pang
 * @Description: 读取文件
 * @receiver t
 * @param fileName
 */
func (t *ReadTxtFile) Read(fileName string) {
	fmt.Println("读取txt文件" + fileName)
}

/**
 * @Author: Jason Pang
 * @Description: 接受访问者类
 * @receiver p
 * @param v
 */
func (t *ReadTxtFile) Accept(v VistorReadFile) {
	v.VistorTxtFile(t)
}

/**
 * @Author: Jason Pang
 * @Description: 访问者，包含对pdf和txt的操作
 */
type VistorReadFile interface {
	VistorPdfFile(p *ReadPdfFile)
	VistorTxtFile(t *ReadTxtFile)
}

/**
 * @Author: Jason Pang
 * @Description: 提取文件类
 */
type ExactFile struct {
}

/**
 * @Author: Jason Pang
 * @Description: 提取pdf文件
 * @receiver e
 * @param p
 */
func (e *ExactFile) VistorPdfFile(p *ReadPdfFile) {
	fmt.Println("提取pdf文件内容")
}

/**
 * @Author: Jason Pang
 * @Description: 提取txt文件
 * @receiver e
 * @param p
 */
func (e *ExactFile) VistorTxtFile(p *ReadTxtFile) {
	fmt.Println("提取txt文件内容")
}

/**
 * @Author: Jason Pang
 * @Description: 压缩文件类
 */
type CompressionFile struct {
}

/**
 * @Author: Jason Pang
 * @Description: 压缩pdf文件
 * @receiver c
 * @param p
 */
func (c *CompressionFile) VistorPdfFile(p *ReadPdfFile) {
	fmt.Println("压缩pdf文件内容")
}

/**
 * @Author: Jason Pang
 * @Description: 压缩txt文件
 * @receiver c
 * @param p
 */
func (c *CompressionFile) VistorTxtFile(p *ReadTxtFile) {
	fmt.Println("压缩txt文件内容")
}

func vistorMain() {
	filesList := []ReadFile{
		&ReadPdfFile{},
		&ReadTxtFile{},
		&ReadPdfFile{},
		&ReadTxtFile{},
	}
	//提取文件
	fmt.Println("--------------------------提取文件")
	extract := ExactFile{}
	for _, f := range filesList {
		f.Accept(&extract)
	}
	//压缩文件
	fmt.Println("--------------------------压缩文件")
	compress := CompressionFile{}
	for _, f := range filesList {
		f.Accept(&compress)
	}
}
