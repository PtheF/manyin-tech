package service

import (
	"github.com/astaxie/beego/logs"
	"go-service/cache"
	"go-service/models"
	"go-service/utils"
)

func AddProduct(prod *models.Product) error {
	prod.ProdId = utils.NextUuid()

	err, avatarName := utils.SaveBase64Img("item-avatar-"+prod.ProdId, "/usr/local/www/assets/img/", prod.ProdAvatar)

	prod.ProdAvatar = "/assets/img/" + avatarName

	err, digitalName := utils.SaveBase64Img("item-digital-"+prod.ProdId, "/usr/local/www/assets/img/", prod.ProdDigital)

	prod.ProdDigital = "/assets/img/" + digitalName

	if err != nil {
		logs.Error("add product service error: %v", err)
		return err
	}

	if err = models.AddProduct(prod); err != nil {
		logs.Error("service > models.AddProduct error: %v", err)
		return err
	}

	return nil
}

func DelProduct(prodId string) error {
	return models.DelProduct(prodId)
}

func GetProductById(prodId string) (models.Product, error) {

	//prod := &models.Product{}

	var prod models.Product

	err := cache.Client("prod", prodId).CacheExp(60 * 60 * 24 * 5).
		DBCallback(func() (interface{}, error) {
			logs.Info("product cache not hit, db callback")
			//if p, err := models.GetProductById(prodId); err != nil {
			//	logs.Error("db not hit")
			//	return nil, errors.New("not hit")
			//}

			p, err := models.GetProductById(prodId)

			if err != nil {
				return nil, err
			}

			prod = p

			logs.Info("inner db callback get prod: %v", prod)

			return prod, nil
		},
		).Query(&prod)

	if err != nil {
		logs.Error("cache error: %v", err)
		return prod, err
	}

	//logs.Info("product service get product by id: id=%v, prod=%v", prodId, prod)

	return prod, nil

}

func GetProductPageByType(currentPage, pageCount, typeId int) (*models.ProductPage, error) {
	//return nil, nil
	return models.GetProductByType(currentPage, pageCount, typeId)
}
