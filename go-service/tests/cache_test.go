package test

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"go-service/cache"
	"go-service/models"
	"testing"
)

func TestCache(t *testing.T) {
	//logs.Info("hello world")

	var addresses []models.Address
	userId := "dfddc9ef211c4907bc2e165930b03329"
	//userId := "dfddc9ef211c4907bc2e165930b03320"

	err := cache.Client("addresses", userId).
		CacheExp(60 * 60 * 24 * 7).
		DBCallback(func() (interface{}, error) {
			if addresses = models.GetAllAddresses(userId); addresses == nil {
				return nil, errors.New("db not hit")
			} else {
				return addresses, nil
			}

		}).Query(&addresses)

	if err != nil {
		logs.Error("cache error: %v ", err)
	} else {
		logs.Info("cache ok")
	}

	for _, address := range addresses {
		logs.Info(address)
	}

	if err = cache.Client("addresses", userId).Remove(); err != nil {
		logs.Error("remove error: %v", err)
	} else {
		logs.Info("remove ok")
	}

}
