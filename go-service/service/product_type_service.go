package service

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"go-service/cache"
	"go-service/models"
)

const (
	ProdTypePrefix = "prodType"
)

func GetProductTypes() []models.ProductType {

	var prodTypes []models.ProductType

	err := cache.Client(ProdTypePrefix, "0").CacheExp(60 * 60 * 24 * 7).
		DBCallback(func() (interface{}, error) {
			if prodTypes = models.GetProductTypes(); prodTypes == nil {
				return nil, errors.New("not hit")
			}

			return prodTypes, nil
		}).Query(&prodTypes)

	if err != nil {
		logs.Error("product types cache client error: %v", err)
		return nil
	}

	return prodTypes

}

func UpdateProdType(prodType models.ProductType) error {
	if err := models.UpdateProductType(prodType); err != nil {
		return err
	}

	_ = cache.Client(ProdTypePrefix, "0").Remove()

	return nil
}
