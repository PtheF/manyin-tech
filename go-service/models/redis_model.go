package models

import (
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {

	server := beego.AppConfig.String("redis.server")
	db, _ := beego.AppConfig.Int("redis.database")
	pw := beego.AppConfig.String("redis.pass")

	maxIdle, _ := beego.AppConfig.Int("redis.maxIdle")
	maxActive, _ := beego.AppConfig.Int("redis.maxActive")
	idleTimeout, _ := beego.AppConfig.Int("idleTimeout")

	err := template.BuildRedTemp().
		Server(server).
		DB(db).
		Auth(pw).
		MaxIdle(maxIdle).
		MaxActive(maxActive).
		IdleTimeout(idleTimeout).Build()

	if err != nil {
		logs.Error("RedisTemplateInitError")
	}

}
