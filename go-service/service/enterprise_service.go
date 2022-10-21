package service

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"go-service/cache"
	"go-service/models"
)

const (
	EnterpriseCachePrefix = "enterprise"
	EnterpriseCacheExpire = 60 * 60 * 24 * 14
)

func AddEnterprise(e *models.EnterprisePO) error {

	if _, err := NewCaptcha(e.OperatorPhone).Verify(e.Captcha); err != nil {
		return err
	}

	if err := models.AddEnterprise(e); err != nil {
		return err
	}

	_ = cache.Client(EnterpriseCachePrefix, e.UserId).Remove()

	return nil
}

func GetEnterpriseByUser(userId string) *models.EnterprisePO {

	enterprise := &models.EnterprisePO{}

	err := cache.Client(EnterpriseCachePrefix, userId).CacheExp(EnterpriseCacheExpire).
		DBCallback(
			func() (interface{}, error) {
				if enterprise = models.GetEnterpriseByUser(userId); enterprise == nil {
					return nil, errors.New("not hit")
				} else {
					return enterprise, nil
				}
			}).Query(enterprise)

	if err != nil {
		logs.Error("get enterprise service error: %v", err)
		return nil
	}

	return enterprise
}
