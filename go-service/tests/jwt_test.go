package test

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt"
	"go-service/utils"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestJwtUtils(t *testing.T) {
	jwtStr, err := utils.NewMapClaimsJwt(func() jwt.MapClaims {
		return jwt.MapClaims{
			"name": "jack",
			"uuid": "201922450108",
		}
	}).Exp(20).Iss("www.chl.com").Sign("secret")

	if err != nil {
		logs.Error("sign jwt error: %v", err)
	} else {
		logs.Info("sign jwt ok: %v", jwtStr)
	}

}
