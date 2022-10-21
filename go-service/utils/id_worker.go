package utils

import (
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

const (
	StartStamp = 1661439413 // 2022-08-25 22:56
	BitOffset  = 32
)

// NextId
/**
  Redis 实现的全局Id生成器
  支持每天 2^32 的id，再多就不行了，当然可以设计成每秒 2^32 个ID
  可以，但没必要
*/
func NextId(key string) int64 {

	now := time.Now().Unix()

	timeStr := time.Unix(now, 0).Format("2006-01-02")

	if count, err := template.OpsForValue().Incr("icr:" + key + ":" + timeStr); err != nil {
		logs.Error("id worker error: %v", err)
		return -1
	} else {
		return ((now - StartStamp) << BitOffset) | int64(count)
	}
}

func NextUuid() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
