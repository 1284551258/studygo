package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwLock     sync.RWMutex
}

/*
1、实现构造函数
2、实现3个方法
	Init(addr string, options ...string) (err error)
	CreateSession(session Session, err error)
	Get(sessionId string) (session Session, err error)
*/

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	//加锁
	m.rwLock.Lock()
	//解锁
	defer m.rwLock.Unlock()

	//用uuid当作SessionId
	uid := uuid.NewV4()
	id := uid.String()
	session = NewMemorySession(id)
	//加入大map
	m.SessionMap[id] = session
	return

}

func (m *MemorySessionMgr) Get(sessionId string) (Session, error) {
	//加锁
	m.rwLock.RLock()
	//解锁
	defer m.rwLock.RUnlock()
	session, ok := m.SessionMap[sessionId]
	if !ok {
		return nil, errors.New("session not exsit in the sessionmgr")
	}
	return session, nil
}
