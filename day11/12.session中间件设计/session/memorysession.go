package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	SessionId string
	Data      map[string]interface{}
	rwLock    sync.RWMutex
}

/*
1、实现构造函数
2、实现这四个方法
Set(key string, value interface{}) error
Get(key string) (interface{}, error)
Del(key string) error
Save() error
*/
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加锁
	m.rwLock.Lock()
	//解锁
	m.rwLock.Unlock()
	m.Data[key] = value
	return
}
func (m *MemorySession) Get(key string) (interface{}, error) {
	//加锁
	m.rwLock.RLock()
	//解锁
	m.rwLock.RUnlock()
	value, ok := m.Data[key]
	if !ok {
		return nil, errors.New("key not exsit in the session")
	}
	return value, nil
}
func (m *MemorySession) Del(key string) (err error) {
	//加锁
	m.rwLock.Lock()
	//解锁
	m.rwLock.Unlock()
	delete(m.Data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}

func NewMemorySession(sessionId string) *MemorySession {
	return &MemorySession{
		SessionId: sessionId,
		Data:      make(map[string]interface{}, 16),
	}
}
