package main

import (
	"github.com/astaxie/beego"
	_ "go-service/models"
	_ "go-service/routers"
)

func main() {
	beego.Run()
}
