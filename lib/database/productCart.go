package libdb

import (
	"errors"
	"fmt"
	//"strings"
	"gorm.io/gorm"

	"berbagi/config"
	"berbagi/models"
)

func GetProductCartByUserId(donorId int) (models.ProductCartGetResponse, error) {
	var productCart []models.ProductCartGetAPI

	res := config.Db.Table("product_carts").Select("product_carts.recipient_id, product_carts.product_package_id, product_carts.quantity").Where(`product_carts.donor_id = ?`, donorId).Find(&productCart)

	if res.Error != nil {
		return models.ProductCartGetResponse{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.ProductCartGetResponse{}, errors.New("No product found in the product cart")
	}

	dictPackage := map[int]bool{}

	for _, v := range productCart {
		dictPackage[int(v.ProductPackageID)] = true
	}

	wantedPackage := []int{}

	for k,_ := range dictPackage {
		wantedPackage = append(wantedPackage, k)
	}

	packageDetails := []models.PackageDetailAPI{}

	joinCondition := ""
	for i := 0; i < len(wantedPackage); i++ {
		joinCondition += fmt.Sprintf("product_package_details.product_package_id = %v", wantedPackage[i])
		if i < len(wantedPackage) - 1 {
			joinCondition += " or "
		}
	}

	productSearch := config.Db.Table("product_package_details").Select("product_package_details.product_package_id, product_package_details.quantity, products.id as product_id, products.price").Joins("left join products on products.id = product_package_details.product_id").Where(joinCondition).Find(&packageDetails)

	if productSearch.Error != nil {
		return models.ProductCartGetResponse{}, productSearch.Error
	}
	response := models.ProductCartGetResponse{}
	response.ItemsBought = productCart
	response.PackageDetail = packageDetails

	return response, nil
}

// This function assumes the userId is still exist. That check should be handled by another auth functionality, not by this function.
func UpdateProductCartByUserId(userCart []models.ProductCart, donorId int)  error {
	err := config.Db.Transaction(func(tx *gorm.DB) error {
		for _, cartItem := range userCart {
			
			// Request body binding done in the controller should already "convert" any integer less than zero to zero
			if cartItem.Quantity == 0 || cartItem.ProductPackageID == 0 {
				continue
			}

			if cartItem.DonorID == cartItem.RecipientID {
				return errors.New("You cant donate to yourself. Please specify different donorID and RecipientID")
			}

			targetCart := models.ProductCart{}

			//Just found about this awesome and convenient method the night before presentation 
			res := tx.Where(models.ProductCart{DonorID: uint(donorId), RecipientID: uint(cartItem.RecipientID), ProductPackageID: cartItem.ProductPackageID}).Assign(models.ProductCart{Quantity: cartItem.Quantity}).FirstOrCreate(&targetCart)

			if res.Error != nil {
				// // Error 1452 means we try to change a child table with invalid parent's table primary key
				// if strings.HasPrefix(res.Error.Error(), "Error 1452") {
				// 	return errors.New(fmt.Sprintf("No product package id %v found in the product package table", cartItem.ProductPackageID))
				// }

				return res.Error
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func DeleteProductCartByUserId(items []models.ProductCartDelAPI, userId int) (error) {
	if len(items) == 0 {
		return errors.New("No item found in delete list. Please specify before deleting")
	}

	deletedCart := models.ProductCart{}

	err := config.Db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			deleteRes := tx.Table("product_carts").Where("donor_id = ? and recipient_id = ? and product_package_id = ?", userId, item.RecipientID, item.ProductPackageID).Unscoped().Delete(&deletedCart)
			
			if deleteRes.Error != nil {
				return deleteRes.Error
			}

			if deleteRes.RowsAffected == 0 {
				return errors.New(fmt.Sprintf("No product package with id %v is found in user's product cart.", item.ProductPackageID))
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}