package cache

import (
	"encoding/json"
	"fmt"
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego/logs"
)

type DBCallbackFunc func() (interface{}, error)

const (
	NilCache          = "NIL_CACHE_PENETRATION" // 缓存穿透存储值
	PenetrationExpire = 60 * 10                 // 缓存穿透过期时间
)

type CahError struct {
	DBNotHit      bool
	Penetration   bool
	CauseError    error
	SetCacheError bool
	JsonError     bool
	CallbackNil   bool
	RemoveError   bool
}

func (e *CahError) Error() string {

	if e.CallbackNil {
		return "db callback nil;"
	}

	if e.DBNotHit {
		return "database not hit;" + e.CauseError.Error()
	}

	if e.Penetration {
		return "nil cache, avoid cache penetration;"
	}

	if e.SetCacheError {
		return "set cache error;" + e.CauseError.Error()
	}

	if e.JsonError {
		return "json error;" + e.CauseError.Error()
	}

	if e.RemoveError {
		return "remove cache error;" + e.CauseError.Error()
	}

	return e.CauseError.Error()
}

type CahClient struct {
	prefix      string
	id          string
	cacheKey    string
	cacheExpire int64
	dbCallback  DBCallbackFunc
}

func Client(prefix string, id string) *CahClient {
	return &CahClient{
		cacheKey:    fmt.Sprintf("cache:%s:%s", prefix, id),
		prefix:      prefix,
		id:          id,
		cacheExpire: 60 * 30,
	}
}

func (c *CahClient) getCache(key string) (string, error) {
	if v, err := template.OpsForValue().Get(key); err != nil {
		logs.Error("cache client get cache error: %v", err)
		return "", err
	} else {
		return v, nil
	}
}

func (c *CahClient) setCacheWithExpire(value interface{}, expire int64) error {

	if value == NilCache {
		_ = template.OpsForValue().SetEx(c.cacheKey, NilCache, expire)
		return nil
	}

	//logs.Info("set cache with expire, value=%v", value)

	if res, err := json.Marshal(value); err != nil {
		logs.Error("json marshal error: %v", err)
		return err
	} else {
		if err := template.OpsForValue().SetEx(c.cacheKey, string(res), expire); err != nil {
			logs.Error("redis set error: %v", err)
			return err
		}

		return nil
	}
}

func (c *CahClient) Remove() *CahError {

	if _, err := template.OpsForKey().Delete(c.cacheKey); err != nil {
		logs.Error("remove cache error: %v", err)
		return &CahError{RemoveError: true, CauseError: err}
	}

	return nil
}

func (c *CahClient) CacheExp(expire int64) *CahClient {
	c.cacheExpire = expire
	return c
}

func (c *CahClient) DBCallback(callback DBCallbackFunc) *CahClient {
	c.dbCallback = callback
	return c
}

func (c *CahClient) Query(v interface{}) *CahError {
	res, err := c.getCache(c.cacheKey)

	if err != nil {
		logs.Info("cache key=%v not hit, call db callback", c.cacheKey)
		if c.dbCallback == nil {
			logs.Error("callback nil")
			return &CahError{CallbackNil: true}
		}

		if dbRes, err := c.dbCallback(); err != nil {
			logs.Info("db not hit, save nil cache")
			_ = c.setCacheWithExpire(NilCache, PenetrationExpire)
			return &CahError{DBNotHit: true, CauseError: err}
		} else {
			if err = c.setCacheWithExpire(dbRes, c.cacheExpire); err != nil {
				logs.Error("set cache with expire error: %v", err)
				return &CahError{SetCacheError: true, CauseError: err}
			}

			return nil
		}

	}

	if res == NilCache {
		return &CahError{Penetration: true}
	}

	if err = json.Unmarshal([]byte(res), v); err != nil {
		return &CahError{JsonError: true, CauseError: err}
	}

	logs.Info("cache key:%v hit", c.cacheKey)

	return nil
}
