package session

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type RedisSession struct {
	SessionId string
	Data      map[string]interface{}
	rwLock    sync.RWMutex
	pool      *redis.Pool
	flag      int
}

const (
	SessionFlagNone = iota
	SessionFlagModify
)

/*
1、实现构造函数
2、实现这四个方法
Set(key string, value interface{}) error
Get(key string) (interface{}, error)
Del(key string) error
Save() error
*/

func NewRedisSession(sessionId string, pool *redis.Pool) *RedisSession {
	return &RedisSession{
		SessionId: sessionId,
		Data:      make(map[string]interface{}, 10),
		pool:      pool,
		flag:      SessionFlagNone,
	}
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.Data[key] = value
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Save() error {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	if r.flag != SessionFlagModify {
		return nil
	}
	//内存中的session的data进行序列化
	data, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.SessionId, string(data))
	if err != nil {
		return err
	}
	r.flag = SessionFlagNone
	return nil
}
func (r *RedisSession) Get(key string) (interface{}, error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	res, ok := r.Data[key]
	if !ok {
		resFromRedis, err := r.loadFromRedis(key)
		if err != nil {
			return nil, err
		}
		return resFromRedis, nil
	}
	return res, nil
}
func (r *RedisSession) loadFromRedis(key string) (interface{}, error) {
	conn := r.pool.Get()

	res, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	data, err := redis.String(res, err)
	if err != nil {
		return nil, err
	}
	//反序列化到内存中
	err = json.Unmarshal([]byte(data), &r.Data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *RedisSession) Del(key string) error {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	delete(r.Data, key)
	return nil
}
