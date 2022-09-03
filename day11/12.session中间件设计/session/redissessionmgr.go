package session

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	//redis地址
	addr string
	//redis密码
	passwd string
	//redis连接池
	pool *redis.Pool
	//锁
	rwLock sync.RWMutex
	//大map
	SessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	//TODO implement me
	//panic("implement me")
	//有传密码
	if len(options) != 0 {
		r.passwd = options[0]
	}
	//创建redis连接池
	r.addr = addr
	r.pool = myRedisPool(r.addr, r.passwd)
	return
}
func myRedisPool(addr, password string) *redis.Pool {

	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			//若有密码，判断
			_, err = conn.Do("AUTH", password)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil

		},
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	//TODO implement me
	//panic("implement me")
	//加锁
	r.rwLock.Lock()
	//解锁
	defer r.rwLock.Unlock()

	//用uuid当作SessionId
	uid := uuid.NewV4()
	id := uid.String()
	session = NewRedisSession(id, r.pool)
	//加入大map
	r.SessionMap[id] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (Session, error) {
	//TODO implement me
	//panic("implement me")
	//加锁
	r.rwLock.RLock()
	//解锁
	defer r.rwLock.RUnlock()
	session, ok := r.SessionMap[sessionId]
	if !ok {
		return nil, errors.New("session not exsit in the sessionmgr")
	}
	return session, nil
}

/*
1、实现构造函数
2、实现3个方法
	Init(addr string, options ...string) (err error)
	CreateSession(session Session, err error)
	Get(sessionId string) (session Session, err error)
*/

func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		SessionMap: make(map[string]Session, 32),
	}
}
