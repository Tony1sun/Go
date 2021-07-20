package session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

// 对象
// MemorySession设计
// 定义MemorySession对象(字段:sessionId、存kv的map，读写锁)
// 构造函数，为了获取对象
// Set()
// Get()
// Del()
// Save()

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	// 设置session，可以先放在内存的map中
	// 批量导入redis,提升性能
	sessionMap map[string]interface{}
	// 读写锁
	rwlock *sync.RWMutex
	// 记录内存中map是否被操作
	flag int
}

// 用常量去定义状态
const (
	// 内存数据没变化
	SessionFlagNone = iota
	SessionFlagModify
)

func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
	return s
}

// session存储到内存中的map
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//设置值
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}

// session存储到redis
func (r *RedisSession) Save() (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 若数据没变，不需要存
	if r.flag != SessionFlagModify {
		return
	}
	// 内存中sessionMap进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取redis连接
	conn := r.pool.Get()
	//保存kv
	_, err = conn.Do("SET", r.sessionId, string(data))
	// 改状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 先判断内存
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exits")
	}
	return
}

func (r *RedisSession) LoadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	// 转字符串
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 取到的东西，反序列化到内存的map
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
