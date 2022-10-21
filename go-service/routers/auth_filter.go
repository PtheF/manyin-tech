package routers

import (
	"errors"
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt"
	"go-service/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var jwtSecret string = beego.AppConfig.String("jwt")
var refreshTokenPrefix string = beego.AppConfig.String("refreshTokenPrefix")

//var authPaths []string = strings.Split(beego.AppConfig.String("authPatterns"), ",")

// refreshToken 不知道有啥用，可能是用于扩展吧
type refreshToken struct {
	stamp string
}

// 获取 Token 的方法，这里暂时用 cookie，可能以后会改成 header
func getToken(c *context.Context) string {
	return c.Input.Header("Authorization")
	//return c.Input.Cookie("token")
}

// 回写 Token 的方法
func rewriteToken(c *context.Context, jwtStr string) {
	c.Output.Header("Access-Control-Expose-Headers", "Authorization")
	c.Output.Header("Authorization", jwtStr)
	//c.Output.Cookie("token", jwtStr)
}

// ReLoginBody 重新登录的json
const ReLoginBody = `{"code":302, "message":"登录过期，请重新登录"}`

func reLogin(c *context.Context) {
	c.Output.Header("Content-Type", "application/json;charset=utf-8")
	_, _ = c.ResponseWriter.Write([]byte(ReLoginBody))
}

// AuthPassBody 认证通过json
const AuthPassBody = `{"code": 200, "message": "ok"}`

func authPass(c *context.Context) {
	c.Output.Header("Content-Type", "application/json;charset=utf-8")
	_, _ = c.ResponseWriter.Write([]byte(AuthPassBody))
}

func RefreshToken(c *context.Context) error {
	//jwtStr := c.Input.Header("Authorization")

	jwtStr := getToken(c)

	//logs.Debug("filter jwt: %v", jwtStr)

	if jwtStr == "" {
		return errors.New("jwt空")
	}

	token, err := jwt.Parse(jwtStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	//logs.Info(token)

	// accessToken 没有任何问题
	if err == nil {

		jwtClaims := token.Claims.(jwt.MapClaims)

		logs.Info("id: %v", jwtClaims["Uuid"])

		// 1. 检查 refreshToken 是否存在
		v, err := template.OpsForValue().Get(refreshTokenPrefix + jwtClaims["Uuid"].(string))

		// refreshToken 过期，重新登录
		if err != nil {
			logs.Error("redis error: %v", err)
			return err
		}

		// 2. 检查 refreshToken 时间戳和 accessToken 时间戳是否吻合
		if jwtClaims["stamp"] != v {
			logs.Warning("access ok but stamp not match, jwt stamp=%v, redis stamp=%v", jwtClaims["stamp"], v)
			return errors.New("时间戳不吻合")
		}

		logs.Info("accessToken ok, let it pass")

		// 3. 放行
		return nil
	}

	// accessToken 有问题
	jwtErr := err.(*jwt.ValidationError)

	// accessToken 过期
	if jwtErr.Errors == 16 {

		logs.Info("accessToken expired, start to refresh")

		jwtClaims := token.Claims.(jwt.MapClaims)

		// 1. 检查 refreshToken 是否存在
		refreshKey := refreshTokenPrefix + jwtClaims["Uuid"].(string)

		//   refreshToken 过期，重新登录 error: redigo: nil returned
		v, err := template.OpsForValue().Get(refreshKey)
		if err != nil {
			logs.Error("redis error: %v", err)
			return err
		}

		// 2. 检查 refreshToken 时间戳和 accessToken 时间戳是否吻合
		if jwtClaims["stamp"] != v {
			logs.Warning("access expired and stamp not match, jwt stamp=%v, redis stamp=%v", jwtClaims["stamp"], v)
			return errors.New("时间戳不吻合")
		}

		// 3. 刷新两个 token，refreshToken 存回 redis，accessToken 放到 responseHeader 中
		newStamp := strconv.FormatInt(time.Now().Unix(), 10)
		err = template.OpsForValue().SetEx(refreshKey, newStamp, 1200)
		if err != nil {
			logs.Error("RedisError: %v", err)
			return errors.New("RedisError")
		}

		at, err := utils.NewMapClaimsJwt(func() jwt.MapClaims {
			return jwtClaims
		}).Exp(600).Stamp(newStamp).Sign(jwtSecret)

		if err != nil {
			return err
		}

		//c.Output.Header("Authorization", at)
		//c.Output.Cookie("token", at)

		rewriteToken(c, at)

		logs.Info("token refresh ok")

		// 4. 放行
		return nil

	} else {
		return jwtErr
	}
}

var JwtAuthFilter = func(c *context.Context) {

	logs.Info("AuthFilter")

	logs.Info("request uri: %v", c.Input.URI())

	//if c.Input.Header("Authorization") == "" {
	if getToken(c) == "" {
		logs.Error("NoJWT")
		reLogin(c)
		return
	} else {
		err := RefreshToken(c)
		if err != nil {
			logs.Error("RefreshTokenError: %v", err)
			reLogin(c)
			return
		}
	}

	// 如果请求接口 == /v1/auth
	if c.Input.URI() == "/v1/auth" {
		authPass(c)
	}
}

var JwtRefreshFilter = func(c *context.Context) {

	logs.Info("RefreshFilter")

	if getToken(c) != "" {
		_ = RefreshToken(c)
	}
}

var success = []byte("SUPPORT OPTIONS")

var corsFunc = func(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Authorization,Admin-Token")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)

	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}

// -------------------------------------------------------------
// 管理员 filter
var adminFilter = func(ctx *context.Context) {
	if ctx.Input.URI() == "/v1/admin/login" {
		return
	}
	adminToken := ctx.Input.Header("Admin-Token")
	if v, _ := template.OpsForKey().Exist("admin:login:" + adminToken); !v {
		//_ = ctx.Output.Body([]byte(ReLoginBody))

		//reLogin(ctx)

		ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
		_, _ = ctx.ResponseWriter.Write([]byte(`{"code": 10302, "message": "登陆过期"}`))
	}
}

func init() {
	/*
	  a) BeforeStatic 静态地址之前
	  b) BeforeRouter 寻找路由之前
	  c) BeforeExec 找到路由之后，开始执行相应的 Controller 之前
	  d) AfterExec 执行完 Controller 逻辑之后执行的过滤器
	  e) FinishRouter 执行完逻辑之后执行的过滤器
	*/
	patterns := strings.Split(beego.AppConfig.String("authPatterns"), ",")

	beego.InsertFilter("*", beego.BeforeExec, JwtRefreshFilter)

	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)

	beego.InsertFilter("/v1/admin/*", beego.BeforeRouter, adminFilter)

	for _, pattern := range patterns {
		beego.InsertFilter(strings.TrimSpace(pattern), beego.BeforeRouter, JwtAuthFilter)
	}
}
