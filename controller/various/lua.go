package various

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client
var script string = `
		local value = redis.call("Get", KEYS[1])
		print("当前值为 " .. value);
		if( value - KEYS[2] >= 0 ) then
			local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
   			print("剩余值为" .. leftStock );
			return leftStock
		else
			print("数量不够，无法扣减");
			return value - KEYS[2]
		end
		return -1
	`
var luaHash string

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "111111", // no password set
		DB:       0,        // use default DB
	})
	luaHash, _ = Client.ScriptLoad(script).Result() //返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值，类似于
}

func useLua() {
	Client.FlushAll()
	//设置初始值
	Client.Set("stock", "10", 0)
	//编写脚本 - 检查数值，是否够用，够用再减，否则返回减掉后的结果
	var luaScript = redis.NewScript(script)
	//执行脚本
	n, err := luaScript.Run(Client, []string{"stock", "6"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

func useLuaHash() {
	n, err := Client.EvalSha(luaHash, []string{"stock", "6"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

func main() {
	//useLua()
	useLuaHash()
}
