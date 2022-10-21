package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-service/cache"
	"go-service/models"
	"strings"
)

const (
	AddressCacheExpire   = 60 * 60 * 24 * 7
	AddressesCachePrefix = "addresses"
)

func AddAddress(address *models.Address) (string, bool) {

	address.Id = strings.Replace(uuid.NewV4().String(), "-", "", -1)

	if _, b := models.AddAddress(address); !b {
		return "", false
	}

	_ = cache.Client(AddressesCachePrefix, address.Uuid).Remove()

	return address.Id, true
}

func DeleteAddress(id string, userId string) bool {
	if models.DeleteAddress(id, userId) {
		_ = cache.Client(AddressesCachePrefix, userId).Remove()
		return true
	}

	return false
}

func SetDefault(id string, userId string) bool {
	if models.SetDefault(id, userId) {
		_ = cache.Client(AddressesCachePrefix, userId).Remove()

		return true
	}

	return false
}

func GetAllAddress(userId string) []models.Address {

	var addresses []models.Address

	err := cache.Client(AddressesCachePrefix, userId).
		CacheExp(AddressCacheExpire).
		DBCallback(func() (interface{}, error) {
			if addresses = models.GetAllAddresses(userId); addresses == nil {
				return nil, errors.New("db not hit")
			} else {
				return addresses, nil
			}
		}).Query(&addresses)

	if err != nil {
		return nil
	}

	return addresses
}
