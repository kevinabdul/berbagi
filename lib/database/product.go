package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
)

func GetProducts(categoryId int) ([]models.ProductAPI, error) {
	var products []models.ProductAPI


	if categoryId == 0 {
		prodSearchRes := config.Db.Table("products").
		Select("products.id, products.name, products.category_id as category_id, products.price").Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found in the product table")
		}
	} else {
		prodSearchRes := config.Db.Table("products").Select("products.id, products.name, products.category_id, products.price").
		Where("category_id = ?", categoryId).Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found for the given category")
		}	

	}
	
	return products, nil
}