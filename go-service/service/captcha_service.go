package service

import (
	"errors"
	"github.com/PtheF/go-redis-template/template"
	"strconv"
	"time"
)

const (
	CaptchaIntervalPrefix = "captcha:interval:"
	CaptchaPrefix         = "captcha:"
)

type CaptchaService struct {
	captcha       string
	key           string
	captchaExpire int64
	interval      int64
}

func NewCaptcha(key string) CaptchaService {
	return CaptchaService{key: key, captchaExpire: 300, interval: 60}
}

func (s CaptchaService) Send() (string, error) {

	// 1. 看是否满足时间间隔
	if ex, _ := template.OpsForKey().Exist(CaptchaIntervalPrefix + s.key); ex == true {
		return "", errors.New(strconv.FormatInt(s.interval/60, 10) + "分钟内不能再次发送")
	}

	// 2. 生成验证码存入 redis
	captcha := strconv.FormatInt((time.Now().Unix()%9000)+1000, 10)
	if err := template.OpsForValue().SetEx(CaptchaPrefix+s.key, captcha, s.captchaExpire); err != nil {
		return "", errors.New("redis错误，存储失败")
	}

	_ = template.OpsForValue().SetEx(CaptchaIntervalPrefix+s.key, "1", s.interval)

	return captcha, nil
}

func (s CaptchaService) Verify(captcha string) (bool, error) {

	if value, err := template.OpsForValue().Get(CaptchaPrefix + s.key); value != captcha || err != nil {
		if err != nil {
			return false, errors.New("验证码过期")
		}

		return false, errors.New("验证码错误")
	}

	return true, nil
}
