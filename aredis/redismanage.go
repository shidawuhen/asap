package aredis

var (
	//
	globalArrRedis map[string]*RedisManager
)
const(
	BASEREDIS = "myRedis"
)

func init() {
	globalArrRedis = make(map[string]*RedisManager)
}

func SetRedis(alias string, redis *RedisManager) {
	globalArrRedis[alias] = redis
}

func GetRedis(alias string) (redis *RedisManager) {
	redis = globalArrRedis[alias]
	return
}
