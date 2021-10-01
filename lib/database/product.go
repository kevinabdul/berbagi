package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
)

func GetProducts(packageI int) ([]models.ProductAPI, error) {
	var products []models.ProductAPI

	if packageI == 0 {
		prodSearchRes := config.Db.Table("products").Select("products.id, products.name, categories.id as category_id, products.price").Joins("left join categories on categories.id = products.category_id").Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found in the product table")
		}
	} else {
		prodSearchRes := config.Db.Table("products").Select("products.id, products.name, categories.id as category_id, products.price").Joins("left join categories on categories.id = products.category_id").Where("categories.id = ?", packageI).Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found for the given cateogory")
		}	

	}
	
	return products, nil
}