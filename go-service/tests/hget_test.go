package test

import (
	"github.com/PtheF/go-redis-template/template"
	_ "github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestHGet(t *testing.T) {
	get, err := template.OpsForHash().Get("nothing", "name")

	logs.Info("hget get=%v, err=%v", get, err)
}
