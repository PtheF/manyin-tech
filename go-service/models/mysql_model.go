package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func init() {
	logs.Info("init orm")

	url := beego.AppConfig.String("mysql.url")

	const driverName = "mysql"

	if url == "" {
		panic("MySQL Url Is Empty")
	}

	var maxIdle, maxOpen int
	var err error
	if maxIdle, err = strconv.Atoi(beego.AppConfig.String("mysql.maxIdle")); err != nil {
		maxIdle = 10
	}

	if maxOpen, err = strconv.Atoi(beego.AppConfig.String("mysql.maxOpen")); err != nil {
		maxOpen = 20
	}

	logs.Info("mysql config:url=%v, maxIdle=%v, maxOpen=%v", url, maxIdle, maxOpen)

	if err = orm.RegisterDriver(driverName, orm.DRMySQL); err != nil {
		logs.Error("RegisterMySQL Error: %v", err)
		panic("RegisterMySQL_ERROR")
	}
	alias := beego.AppConfig.String("mysql.alias")
	if alias == "" {
		alias = "default"
	}
	if err = orm.RegisterDataBase(alias, driverName, url, maxIdle, maxOpen); err != nil {
		logs.Error("RegisterDatabase Error: %v", err)
		panic("RegisterDatabase_Error")
	}

	logs.Info("ORM_INIT_FINISHED")
}
