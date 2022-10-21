package test

import (
	"github.com/astaxie/beego/logs"
	"go-service/utils"
	"testing"
	"time"
)

func TestIdWorker(t *testing.T) {
	logs.Info(utils.NextId("order"))
	logs.Info(utils.NextId("order"))

	logs.Info(utils.NextId("user"))
	logs.Info(utils.NextId("user"))

	time.Sleep(10)

	logs.Info(utils.NextId("order"))
}
