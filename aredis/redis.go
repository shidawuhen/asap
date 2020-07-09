package aredis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

const (
	// DEFAULT_POOLSIZE 默认连接池大小
	DEFAULT_POOLSIZE = 5
	DEFAULT_TIMEOUT  = time.Duration(1000) * time.Millisecond
)

// RedisManager 管理长连接池的redis管理器
// 操作失败自动重连
type RedisManager struct {
	Name           string
	host           string
	auth           string
	redisPool      chan redis.Conn
	connectTimeout time.Duration
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

// NewRedisManager 新建redis管理器
func NewRedisManager(name, host, auth string, poolsize int, timeout ...int) (p *RedisManager, err error) {
	p = &RedisManager{}
	if poolsize == 0 {
		poolsize = DEFAULT_POOLSIZE
	}
	p.Name = name
	p.host = host
	p.auth = auth
	p.redisPool = make(chan redis.Conn, poolsize)

	p.connectTimeout = DEFAULT_TIMEOUT
	p.readTimeout = DEFAULT_TIMEOUT
	p.writeTimeout = DEFAULT_TIMEOUT
	tlen := len(timeout)
	if tlen > 0 {
		p.connectTimeout = time.Duration(timeout[0]) * time.Millisecond
	}
	if tlen > 1 {
		p.readTimeout = time.Duration(timeout[1]) * time.Millisecond
	}
	if tlen > 2 {
		p.writeTimeout = time.Duration(timeout[2]) * time.Millisecond
	}
	for i := 0; i < poolsize; i++ {
		conn, err := p.dialRedis()
		if err != nil {
			return nil, err
		}
		p.redisPool <- conn
	}

	return
}

// InfoRM redis manager info
func (w *RedisManager) InfoRM() string {

	return fmt.Sprintf("pool-len:%d, auth:%s, host:%s, connectTimeout:%v, readTimeout:%v, writeTimeout:%v",
		len(w.redisPool), w.auth, w.host, w.connectTimeout, w.readTimeout, w.writeTimeout)
}

func (w *RedisManager) dialRedis() (conn redis.Conn, err error) {
	conn, err = redis.DialTimeout("tcp", w.host, w.connectTimeout, w.readTimeout, w.writeTimeout)
	if err != nil {
		return nil, err
	}
	if w.auth != "" {
		_, err = conn.Do("AUTH", w.auth)
	}
	return
}

// Do 执行闭包函数
func (w *RedisManager) Do(action func(conn redis.Conn) (interface{}, error)) (reply interface{}, err error) {
	return w.redialDo(action)
}

func (w *RedisManager) redialDo(action func(conn redis.Conn) (reply interface{}, err error)) (reply interface{}, err error) {
	conn := w.getConn()
	count := 0

start:
	count++
	if count > 3 {
		w.putConn(conn)
		return
	}
	reply, err = action(conn)
	if err != nil {
		goto fail_redial
	}
	w.putConn(conn)
	return

fail_redial:
	conn.Close()
	var nerr error
	var nconn redis.Conn
	if nconn, nerr = w.dialRedis(); nerr != nil {
		err = nerr
		goto start
	}
	conn = nconn
	goto start
}

func (w *RedisManager) getConn() redis.Conn {
	return <-w.redisPool
}

func (w *RedisManager) putConn(conn redis.Conn) {
	w.redisPool <- conn
}

// ----------- string -----------------

// Set 设置key 和value
// redis commands
func (w *RedisManager) Set(key string, value interface{}) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SET", key, value)
	}
	return w.Do(action)
}

// SetEx 设置key 和value，并且设置超时(s)
// redis commands
func (w *RedisManager) SetEx(key string, value interface{}, expire int64) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SETEX", key, expire, value)
	}
	return w.Do(action)
}

// SetExNx set if not exist, with an expire time(s)
func (w *RedisManager) SetExNx(key string, value interface{}, expire int64) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SET", key, value, "EX", expire, "NX")
	}
	return w.Do(action)
}

// PSetEx 设置key 和value，并且设置超时(ms)
func (w *RedisManager) PSetEx(key string, value interface{}, expire int64) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("PSETEX", key, expire, value)
	}
	return w.Do(action)
}

// MSet set多重的k-v
func (w *RedisManager) MSet(param map[string]interface{}) (reply interface{}, err error) {

	args := []interface{}{}
	for k, v := range param {
		args = append(args, k, v)
	}
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("MSET", args...)
	}

	reply, err = w.Do(action)
	return

}

// Get 获取value
func (w *RedisManager) Get(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("GET", key)
	}
	return w.Do(action)
}

// MGet 一次获取多个key的值
func (w *RedisManager) MGet(keys []string) (result map[string]string, err error) {

	args := []interface{}{}
	for _, v := range keys {
		args = append(args, v)
	}

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("MGET", args...)
	}
	reply, err := w.Do(action)
	if err != nil {
		return
	}

	result, err = w.replyToMap(reply, keys)
	return
}

// Exists 检测 key 是否存在.
func (w *RedisManager) Exists(key string) (is bool, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("EXISTS", key)
	}

	return redis.Bool(w.Do(action))
}

// Del 删除 key.
func (w *RedisManager) Del(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("DEL", key)
	}

	return w.Do(action)
}

// Incr key 自加.
func (w *RedisManager) Incr(key string) (num int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("INCR", key)
	}
	num, err = redis.Int64(w.Do(action))
	return
}

// Incrby 原子加
func (w *RedisManager) Incrby(key string, add int64) (num int64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("INCRBY", key, add)
	}

	num, err = redis.Int64(w.Do(action))
	return
}

// Decr key 自减.
func (w *RedisManager) Decr(key string) (num int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("DECR", key)
	}
	num, err = redis.Int64(w.Do(action))
	return
}

// GetString 获取数据，转化为string
func (w *RedisManager) GetString(key string) (value string, err error) {
	reply, err := w.Get(key)
	if err != nil {
		return
	}
	if reply == nil {
		return
	}

	return redis.String(reply, err)
}

// GetInt64 获取数据，转化为int64
func (w *RedisManager) GetInt64(key string) (value int64, err error) {
	val, err := w.Get(key)
	if err != nil {
		return
	}
	if val == nil {
		return
	}

	return redis.Int64(val, err)
}

// GetFloat 获取数据，转化为float64
func (w *RedisManager) GetFloat(key string) (value float64, err error) {
	return redis.Float64(w.Get(key))
}

// ----------- hash -----------------

// Hincrby 给hash的key加数
func (w *RedisManager) Hincrby(key string, subkey string, inc int64) (num int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HINCRBY", key, subkey, inc)
	}
	num, err = redis.Int64(w.Do(action))
	return
}

// Hset 设置hash表的subkey值
func (w *RedisManager) Hsetnx(key string, subkey string, value interface{}) (reply int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HSETNX", key, subkey, value)
	}
	reply, err = redis.Int64(w.Do(action))
	return
}

// Hset 设置hash表的subkey值
func (w *RedisManager) Hset(key string, subkey string, value interface{}) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HSET", key, subkey, value)
	}
	return w.Do(action)
}

// Hmset hash的multiset
func (w *RedisManager) Hmset(key string, param map[string]interface{}) (reply interface{}, err error) {

	args := []interface{}{key}
	for k, v := range param {
		args = append(args, k, v)
	}
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HMSET", args...)
	}

	reply, err = w.Do(action)
	return
}

// Hmget hash multiget
func (w *RedisManager) Hmget(key string, param []string) (result map[string]string, err error) {

	args := []interface{}{key}
	for _, v := range param {
		args = append(args, v)
	}
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HMGET", args...)
	}

	reply, err := w.Do(action)
	if err != nil {
		return
	}

	result, err = w.replyToMap(reply, param)

	return
}

// Hget 获取hash表的subkey值
func (w *RedisManager) Hget(key string, subkey string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HGET", key, subkey)
	}
	return w.Do(action)
}

// Hexists 查询一个键值是否存在
func (w *RedisManager) Hexists(key string, subkey string) (result bool, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HEXISTS", key, subkey)
	}
	reply, err := w.Do(action)
	if err != nil {
		return
	}
	if reply == nil {
		return
	}
	res := fmt.Sprintf("%v", reply)
	if res == "1" {
		result = true
	}
	return
}

// Hkeys 获取hash表里key的所有field
func (w *RedisManager) Hkeys(key string) (result []string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HKEYS", key)
	}
	reply, err := w.Do(action)
	result, err = redis.Strings(reply, err)
	if err != nil {
		return
	}

	return
}

// Pttl 获取过期时间，单位：毫秒
func (w *RedisManager) Pttl(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("PTTL", key)
	}
	return w.Do(action)
}

// Ttl ttl 获取过期时间，单位：秒
func (w *RedisManager) Ttl(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("TTL", key)
	}
	return w.Do(action)
}

// PTTL 获取过期时间，单位：毫秒
func (w *RedisManager) PTTL(key string) (pttl int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("PTTL", key)
	}
	return redis.Int64(w.Do(action))
}

// TTL TTL 获取过期时间
func (w *RedisManager) TTL(key string) (ttl int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("TTL", key)
	}
	return redis.Int64(w.Do(action))
}

// HGetAll 获取hash表的所有记录
func (w *RedisManager) HGetAll(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HGETALL", key)
	}
	return w.Do(action)
}

// HgetAll 返回hash所有记录，以k-v方式输出
func (w *RedisManager) HgetAll(key string) (result map[string]string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HGETALL", key)
	}
	reply, err := w.Do(action)
	replyList, err := redis.Strings(reply, err)
	if err != nil {
		return
	}
	var k string
	result = make(map[string]string)
	for i, cur := range replyList {
		if i%2 == 0 {
			k = cur
			continue
		}
		result[k] = cur
	}
	return
}

// HDel 删除 hash 表的 subkey 值.
func (w *RedisManager) HDel(key string, subkey string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HDEL", key, subkey)
	}

	return w.Do(action)
}

// HExists 查询 hash 表中的 subkey 值是否存在.
func (w *RedisManager) HExists(key string, subkey string) (is bool, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HEXISTS", key, subkey)
	}

	return redis.Bool(w.Do(action))
}

// GetHString 获取 HGET 的 string.
func (w *RedisManager) GetHString(key string, subkey string) (value string, err error) {

	return redis.String(w.Hget(key, subkey))
}

// GetHInt64 获取 HGET 的 int.
func (w *RedisManager) GetHInt64(key string, subkey string) (value int64, err error) {
	val, err := w.Hget(key, subkey)
	if err == redis.ErrNil {
		err = nil
		return
	}
	if err != nil {
		return
	}
	if val == nil {
		return
	}

	return redis.Int64(val, err)
}

// ----------- set -----------------

// Sismember 查询是否集合成员
func (w *RedisManager) Sismember(key string, member interface{}) (res bool, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SISMEMBER", key, member)
	}
	return redis.Bool(w.Do(action))
}

// Sadd 添加集合成员
func (w *RedisManager) Sadd(key string, member ...interface{}) (reply interface{}, err error) {
	args := []interface{}{key}
	for _, m := range member {
		args = append(args, m)
	}
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SADD", args...)
	}
	return w.Do(action)
}

// Scard 获取集合元素个数
func (w *RedisManager) Scard(key string) (num int64, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SCARD", key)
	}
	reply, err := w.Do(action)
	if err != nil || reply == nil {
		return
	}
	num, err = redis.Int64(reply, err)
	return
}

// Smembers 返回集合所有成员
func (w *RedisManager) Smembers(key string) (member []string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SMEMBERS", key)
	}
	reply, err := w.Do(action)
	if err != nil {
		return
	}
	if reply == nil {
		return
	}
	member, err = redis.Strings(reply, err)
	return
}

// ----------- system ---------------

// Expire 超时
func (w *RedisManager) Expire(key string, ts int64) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("EXPIRE", key, ts)
	}
	return w.Do(action)
}

// Info 查看redis系统情况
func (w *RedisManager) Info() (stats []string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("INFO")
	}
	content, err := redis.Bytes(w.Do(action))
	stats = strings.Split(string(content), "\n")
	return

}

// ----------- list ---------------

// Lpush 左边push
func (w *RedisManager) Lpush(key string, content interface{}) (err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LPUSH", key, content)
	}
	_, err = w.Do(action)
	return
}

// LpushMulti 批量lpush
func (w *RedisManager) LpushMulti(key string, contents ...interface{}) (reply interface{}, err error) {
	args := []interface{}{key}
	args = append(args, contents...)
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LPUSH", args...)
	}
	return w.Do(action)
}

// Rpop 右边pop
func (w *RedisManager) Rpop(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("RPOP", key)
	}
	reply, err = w.Do(action)
	return
}

// Brpop 右边阻塞 pop
func (w *RedisManager) Brpop(key string, timeout int) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("BRPOP", key, timeout)
	}
	reply, err = w.Do(action)
	return
}

// Llen 长度
func (w *RedisManager) Llen(key string) (length int, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LLEN", key)
	}
	reply, err := w.Do(action)
	if reply == nil {
		err = fmt.Errorf("reply eq nil")
		return
	}
	_, ok := reply.(int64)
	if !ok {
		err = fmt.Errorf("invalid valid")
		return
	}
	length = int(reply.(int64))
	return
}

// Lpop len左边push
func (w *RedisManager) Lpop(key string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LPOP", key)
	}
	reply, err = w.Do(action)
	return
}

// Lrange 获取列表一段
func (w *RedisManager) Lrange(key string, start, end int) (reply []string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LRANGE", key, start, end)
	}
	reply, err = redis.Strings(w.Do(action))
	return
}

// Ltrim 删除区间外的所有元素, 保留区间内的元素
func (w *RedisManager) Ltrim(key string, start, end int) (err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LTRIM", key, start, end)
	}
	_, err = w.Do(action)
	return
}

// Lindex 获取指定的一个元素
func (w *RedisManager) Lindex(key string, index int) (reply string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LINDEX", key, index)
	}
	reply, err = redis.String(w.Do(action))
	return
}

// ----------- zset -----------------

// Zincrby _
func (w *RedisManager) Zincrby(key string, name string, add int64) (num int64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZINCRBY", key, add, name)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	num, err = redis.Int64(reply, err)
	return
}

// Zrank 排名，rank = -1 表示不在zset中
func (w *RedisManager) Zrank(key string, name string) (rank int64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZRANK", key, name)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		rank = -1
		return
	}
	rank, err = redis.Int64(reply, err)
	return
}

// Zrevrank 反向排名，rank = -1 表示不在zset中
func (w *RedisManager) Zrevrank(key string, name string) (rank int64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZREVRANK", key, name)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		rank = -1
		return
	}
	rank, err = redis.Int64(reply, err)
	return
}

// Zrem zset rem
func (w *RedisManager) Zrem(key string, member string) (del int, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZREM", key, member)
	}

	reply, err := w.Do(action)
	if err != nil {
		return
	}

	del, err = redis.Int(reply, err)
	return
}

// Zadd zset add member
func (w *RedisManager) Zadd(key string, member string, score float64) (add int, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZADD", key, score, member)
	}

	reply, err := w.Do(action)
	if err != nil {
		return
	}

	add, err = redis.Int(reply, err)
	return
}

// Zscore key不存在或member不存在，都返回("", err!=nil)
func (w *RedisManager) Zscore(key string, member string) (score string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZSCORE", key, member)
	}

	reply, err := w.Do(action)
	if err != nil {
		return
	}

	score, err = redis.String(reply, err)
	return
}

// Zrange 排名列表
func (w *RedisManager) Zrange(key string, start, end int) (rankList []string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZRANGE", key, start, end)
	}
	rankList, err = redis.Strings(w.Do(action))
	return
}

// Zrevrange 反向排名列表
func (w *RedisManager) Zrevrange(key string, start, end int) (rankList []string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZREVRANGE", key, start, end)
	}
	rankList, err = redis.Strings(w.Do(action))
	return
}

// ZrangeWithScore 排名列表 with score
func (w *RedisManager) ZrangeWithScore(key string, start, end int) (sortedKeys []string, result map[string]float64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZRANGE", key, start, end, "WITHSCORES")
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}

	replyList, err := redis.Strings(reply, err)
	if err != nil {
		return
	}

	if len(replyList)%2 != 0 {
		err = fmt.Errorf("invalid reply")
		return
	}

	sortedKeys = make([]string, 0)
	result = make(map[string]float64)
	for i := 0; i < len(replyList); i += 2 {
		key := replyList[i]
		value, err := strconv.ParseFloat(replyList[i+1], 64)
		if err != nil {
			return nil, nil, err
		}
		result[key] = value
		sortedKeys = append(sortedKeys, key)
	}
	return
}

// ZrevrangeWithScore 反向排名列表 with score
func (w *RedisManager) ZrevrangeWithScore(key string, start, end int) (result map[string]float64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("ZREVRANGE", key, start, end, "WITHSCORES")
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}

	replyList, err := redis.Strings(reply, err)
	if err != nil {
		return
	}

	if len(replyList)%2 != 0 {
		err = fmt.Errorf("invalid reply")
		return
	}

	result = make(map[string]float64)
	for i := 0; i < len(replyList); i += 2 {
		key := replyList[i]
		value, err := strconv.ParseFloat(replyList[i+1], 64)
		if err != nil {
			return nil, err
		}
		result[key] = value
	}
	return
}

// ----------- other ----------------
func (w *RedisManager) replyToMap(reply interface{}, param []string) (result map[string]string, err error) {

	replyList, err := redis.Strings(reply, err)
	if err != nil {
		return
	}
	result = make(map[string]string)
	if len(replyList) != len(param) {
		err = fmt.Errorf("RedisManager: hmget reply list not match ")
		return
	}
	for i, cur := range replyList {
		result[param[i]] = cur
	}

	return
}

// ---------hlen------------------

// Hlen _
func (w *RedisManager) Hlen(key string) (length int64, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("HLEN", key)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		length = -1
		return
	}
	length, err = redis.Int64(reply, err)
	return
}

// Rename 重命名 key.
func (w *RedisManager) Rename(old, name string) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("RENAME", old, name)
	}
	reply, err = w.Do(action)
	return
}

// ----------- keys ----------------
// Keys _
func (w *RedisManager) Keys(pattern string) (keys []string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("keys", pattern)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		return
	}
	keys, err = redis.Strings(reply, err)
	return
}

// ----------- type ----------------
// Keys _
func (w *RedisManager) Type(key string) (result string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("type", key)
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		return
	}
	result, err = redis.String(reply, err)
	return
}

// ----------- bitmap ----------------
// Setbit ...
func (w *RedisManager) SetBit(key string, offset, value uint32) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("setbit", key, offset, value)
	}
	return w.Do(action)
}

// Getbit ...
func (w *RedisManager) GetBit(key string, offset uint32) (reply interface{}, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("getbit", key, offset)
	}
	return w.Do(action)
}

// ----------- server ----------------
// Time 返回当前服务器时间
func (w *RedisManager) Time() (replyList []string, err error) {
	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("TIME")
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}

	replyList, err = redis.Strings(reply, err)
	if err != nil {
		return
	}

	if len(replyList) != 2 {
		err = fmt.Errorf("invalid reply")
		return
	}

	return
}

// Ping
func (w *RedisManager) Ping() (result string, err error) {

	action := func(conn redis.Conn) (interface{}, error) {
		return conn.Do("PING")
	}
	reply, err := w.Do(action)

	if err != nil {
		return
	}
	if reply == nil {
		return
	}
	result, err = redis.String(reply, err)
	return
}
