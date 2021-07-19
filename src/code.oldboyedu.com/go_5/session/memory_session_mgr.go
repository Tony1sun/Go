package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr设计
// 定义MemorySessionMgr对象(字段：存放所有session的map，读写锁)
// 构造函数
// Init()
// CreateSession()
// GetSession()

// 定义对象
type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建session
func (s *MemorySessionMgr) CreatSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Lock()
	// 用uuid作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	// 转string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)
	// 加入到大map
	s.SessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.SessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
