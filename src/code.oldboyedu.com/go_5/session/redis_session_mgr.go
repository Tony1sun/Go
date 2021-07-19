package session

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)

// RedisSessionMgr
// 定义RedisSessionMgr对象(字段：存放所有session的map，读写锁)
// 构造函数
// Init()
// CreateSession()
// GetSession()

type RedisSessionMgr struct {
	// redis地址
	addr string
	// 密码
	passwd string
	// 连接池
	pool *redis.Pool
	// 锁
	rwlock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}
