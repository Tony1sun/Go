package session

// 定义管理者，管理所有session
type SessionMgr interface {
	// 初始化
	Init(addr string, options ...string) (err error)
	CreatSession() (Session Session, err error)
	Get(sessionId string) (Session Session, err error)
}
