package framework

import (
	"net/http"
	"sync"

	"fmt"
)

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

//Context，用于处理请求的输入和输出数据
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

//将请求返回值写入http.Responsewriter中
func (c *Context) String(format string, data ...interface{}) {
	fmt.Fprintf(c.Writer, format, data...)
	return
}

//核心结构，存放路由规则和使用pool获取与释放Context，减少GC
type Engine struct {
	pool   sync.Pool
	router map[string]map[string]HandlersChain
}

//Engine初始化
func New() *Engine {
	fmt.Println("start")
	engine := &Engine{}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	engine.router = make(map[string]map[string]HandlersChain)
	return engine
}

func (engine *Engine) allocateContext() *Context {
	return &Context{}
}

//请求过来时http包会调用该函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	//1.context初始化
	c.Writer = w
	c.Request = req
	//2.真正的处理请求逻辑
	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	//从router中找到对应的方法并执行，如果不存在，则直接返回
	routers, ok := engine.router[httpMethod]
	if ok {
		handles, ok := routers[rPath]
		if ok {
			for _, handle := range handles {
				handle(c)
			}
			return
		}
	}
	c.String("%s", httpMethod+" "+rPath+" doesn't exist")
	return
}

//将路由添加到router中，没有并发操作，所以不加锁
func (engine *Engine) AddRoute(method, path string, handlers ...HandlerFunc) {
	//1.判断该http方法是否存在
	_, ok := engine.router[method]
	if !ok {
		engine.router[method] = make(map[string]HandlersChain)
	}
	//2.判断该路径是否存在，如果不存在则插入，如果存在，则不处理
	_, ok = engine.router[method][path]
	if !ok {
		engine.router[method][path] = handlers
	}
	fmt.Println(engine.router)
}

//运行服务，监听请求
func (engine *Engine) Run(address string) (err error) {
	err = http.ListenAndServe(address, engine)
	return
}
