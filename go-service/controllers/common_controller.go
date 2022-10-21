package controllers

import (
	"encoding/json"
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt"
	"go-service/models"
	"go-service/utils"
	"strconv"
	"time"
)

type ControllerResp struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResp(status int, code int, msg string, data interface{}) *ControllerResp {
	return &ControllerResp{
		Status:  status,
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type BaseController struct {
	beego.Controller
}

var (
	JwtSecret  = beego.AppConfig.String("jwt")
	RefreshKey = beego.AppConfig.String("refreshTokenPrefix")
)

var (
	ParamIllegal = &ControllerResp{200, 101, "请求参数错误", nil}

	DatabaseError = &ControllerResp{200, 102, "数据库挂了", nil}

	UserLoginError = &ControllerResp{200, 103, "密码错误或用户不存在", nil}

	CaptchaError = &ControllerResp{200, 104, "验证码错误", nil}

	UserExist = &ControllerResp{200, 201, "用户已存在", nil}

	VerifyError = &ControllerResp{200, 202, "用户认证失败", nil}

	ServerError = &ControllerResp{200, 500, "服务器错误", nil}
)

func (c *BaseController) Failed(msg string) {
	c.Resp(&ControllerResp{Status: 200, Code: 600, Message: msg, Data: nil})
}

func (c *BaseController) Resp(e *ControllerResp) {

	c.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	// 这一步就是那个请求状态码，200 304 503 啥的，这里暂且这么写，就是 ControllerResp 里面可以控制状态码。
	// 但是这里目前先都弄成200
	// 前后端交互的时候关于这个状态码怎么设置一直有分歧，有人说应该全都返回 200 ok，然后根据返回的 code 进行判断
	// 还有人说直接返回状态码得了呗，弄什么 code
	// 这里就都弄上，即返回 code 也设置 status。
	// 只是设置 status 确实有点麻烦，所以先这么写上，但是 status 都写死200
	c.Ctx.ResponseWriter.WriteHeader(e.Status)

	c.Data["json"] = e

	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Success(msg string) {
	c.Resp(&ControllerResp{Code: 200, Status: 200, Message: msg, Data: nil})
}

func (c *BaseController) ReturnData(data interface{}) {
	c.Resp(&ControllerResp{Status: 200, Code: 200, Message: "ok", Data: data})
}

func (c *BaseController) FormString(key string) string {
	return c.GetString(key)
}

func (c *BaseController) RestString(key string) string {
	return c.Ctx.Input.Param(":" + key)
}

func (c *BaseController) QueryString(key string) string {
	return c.Ctx.Input.Query(key)
}

func (c *BaseController) GetJson(t interface{}) {
	r := c.Ctx.Input.RequestBody

	_ = json.Unmarshal(r, t)
}

func (c *BaseController) SetJWT(jwt string) {
	c.Ctx.Output.Header("Access-Control-Expose-Headers", "Authorization")
	c.Ctx.Output.Header("Authorization", jwt)

	//c.Ctx.Output.Cookie("token", jwt)
}

func (c *BaseController) GetJwt() string {
	return c.Ctx.Input.Header("Authorization")
	//return c.Ctx.Input.Cookie("token")
}

func (c *BaseController) GenJwt(userDTO *models.UserDTO) (string, error) {

	newStamp := strconv.FormatInt(time.Now().Unix(), 10)

	sign, err := utils.NewMapClaimsJwt(func() jwt.MapClaims {
		return jwt.MapClaims{
			"Uuid":    userDTO.Uuid,
			"Email":   userDTO.Email,
			"Phone":   userDTO.Phone,
			"RegDate": userDTO.RegDate,
		}
	}).Stamp(newStamp).Exp(600).Iss("www.heuet.edu.cn").Sign(JwtSecret)

	if err != nil {
		return "", err
	}

	refreshKey := RefreshKey + userDTO.Uuid
	err = template.OpsForValue().SetEx(refreshKey, newStamp, 1200)

	if err != nil {
		return "", err
	}

	return sign, nil
}

func (c *BaseController) GetUser() *models.UserDTO {
	jwtStr := c.GetJwt()

	token, _ := jwt.Parse(jwtStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	claims := token.Claims.(jwt.MapClaims)

	logs.Info("GetUser: %v", claims)

	return &models.UserDTO{
		Uuid:    claims["Uuid"].(string),
		Phone:   claims["Phone"].(string),
		Email:   claims["Email"].(string),
		RegDate: claims["RegDate"].(string),
	}
}
