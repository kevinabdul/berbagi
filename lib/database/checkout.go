package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"time"
	"fmt"
	"errors"
	"strings"
	"gorm.io/gorm"
	"strconv"
)

func GetCheckoutByUserId(userId int) (models.CheckoutGetResponse, error){
	var productCart []models.GiftAPI
	res := config.Db.Table("product_carts").
	Select("product_carts.recipient_id, product_carts.product_package_id, product_carts.quantity").
	Where(`product_carts.donor_id = ?`, userId).Find(&productCart)

	if res.Error != nil {
		return models.CheckoutGetResponse{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.CheckoutGetResponse{}, errors.New("no product_package_id found in user's product_carts")
	}

	dictPackage := map[int]bool{}
	dictGift := map[int][]models.GiftAPI{}

	for _, v := range productCart {
		dictPackage[int(v.ProductPackageID)] = true
		dictGift[int(v.RecipientID)] = append(dictGift[int(v.RecipientID)], v)
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

	productSearch := config.Db.Table("product_package_details").
	Select("product_package_details.product_package_id, product_package_details.quantity, products.id as product_id, products.price").
	Joins("left join products on products.id = product_package_details.product_id").Where(joinCondition).Find(&packageDetails)

	if productSearch.Error != nil {
		return models.CheckoutGetResponse{}, productSearch.Error
	}

	dictPackageDetail := map[int][]models.PackageDetailAPI{}

	for _, v := range packageDetails {
		dictPackageDetail[int(v.ProductPackageID)] = append(dictPackageDetail[int(v.ProductPackageID)], 
		models.PackageDetailAPI{ProductID: uint(v.ProductID), Quantity: int(v.Quantity), Price: v.Price})
	}

	packageLists := []models.PackageListAPI{}

	for k, v := range dictPackageDetail {
		packageList := models.PackageListAPI{}
		packageList.ProductPackageID = uint(k)
		packageList.Details = v
		packageLists = append(packageLists, packageList)
	}

	response := models.CheckoutGetResponse{}
	recipientGifts := []models.RecipientGift{}

	for k, v := range dictGift {
		recipientGifts = append(recipientGifts, models.RecipientGift{RecipientID: uint(k), Gifts: v})
	}

	response.Recipients = recipientGifts
	response.PackageList = packageLists

	paymentMethods := []models.PaymentMethod{}

	if err := config.Db.Table("payment_methods").Find(&paymentMethods).Error; err != nil {
		return models.CheckoutGetResponse{}, err
	}

	response.PaymentMethods = paymentMethods

	return response, nil
}

func AddCheckoutByUserId(payment models.PaymentMethod, donorId int) (models.TransactionAPI, error) {
	var productCart []models.GiftAPI
	findCartRes := config.Db.Table("product_carts").
	Select("product_carts.recipient_id, product_carts.product_package_id, product_carts.quantity").
	Where(`product_carts.donor_id = ?`, donorId).Find(&productCart)

	if findCartRes.Error != nil {
		return models.TransactionAPI{}, findCartRes.Error
	}

	if findCartRes.RowsAffected == 0 {
		return models.TransactionAPI{}, errors.New("no product_package found in the cart. add product_package first before checking out")
	}

	/* transactionAPI, transactionCreation, transaction, and transactionDetail here means:
	   the display struct used for response, query result from gorm method, model, and model,
	   not the mysql transaction statement executed in the function below */
	transactionAPI := models.TransactionAPI{}
	paymentMethod := models.PaymentMethod{}

	var transactionCreation *gorm.DB

	err := config.Db.Transaction(func(tx *gorm.DB) error {

		deletedCart := models.ProductCart{}

		deleteRes := tx.Table("product_carts").Where("donor_id = ?", donorId).Unscoped().Delete(&deletedCart)

		if deleteRes.Error != nil {
			return deleteRes.Error
		}

		if deleteRes.RowsAffected == 0 {
			return errors.New("failed to delete items in donor's cart.")
		}

		packageQuantity := map[uint]int{}
		packagetarget := ""
		dictRecipient := map[uint][]models.GiftAPI{}

		for _, v := range productCart {
			if _, ok := packageQuantity[v.ProductPackageID]; !ok {
				packageQuantity[v.ProductPackageID] = v.Quantity
			} else {
				packageQuantity[v.ProductPackageID] += v.Quantity
			}

			if _, ok := dictRecipient[v.RecipientID]; !ok {
				dictRecipient[v.RecipientID] = append(dictRecipient[v.RecipientID], v)
			}
		}

		for k, _ := range packageQuantity {
			packagetarget += strconv.Itoa(int(k)) + ","
		}

		packagetarget = "(" + packagetarget[0: len(packagetarget) - 1] + ")"

		packagePrices := []models.PackagePrice{}
		packagePriceMap := map[uint]int{}

		prodDetailRes := config.Db.Table("product_package_details").
		Select("product_package_details.product_package_id, sum(products.price) as price").
		Joins("left join products on products.id = product_package_details.product_id").Group("product_package_details.product_package_id").
		Having(fmt.Sprintf("product_package_id in %s", packagetarget)).Find(&packagePrices)
		
		if prodDetailRes.Error != nil {
			return prodDetailRes.Error
		}

		var total int

		for _, v := range packagePrices {
			total += v.Price * int(packageQuantity[v.ProductPackageID])
			packagePriceMap[v.ProductPackageID] = v.Price
		}

		invoiceId := fmt.Sprintf("BERBAGI.DONOR.%03v.%v", donorId, time.Now().String()[0:19])

		transaction := models.Transaction{}
		transaction.DonorID = uint(donorId)
		transaction.InvoiceID = invoiceId
		transaction.PaymentMethodID = payment.ID
		transaction.Total = total

		transactionCreation = tx.Create(&transaction)

		if transactionCreation.Error != nil {
			if strings.HasPrefix(transactionCreation.Error.Error(), "Error 1452") {
				return errors.New(fmt.Sprintf("no payment_method_id '%v' found in the payment method table", payment.ID))
			}
			return transactionCreation.Error
		}

		transactionDetail := models.TransactionDetail{}
		transactionDetail.InvoiceID = invoiceId

		for _, cartItem := range productCart {
			transactionDetail.RecipientID = cartItem.RecipientID
			transactionDetail.ProductPackageID = cartItem.ProductPackageID
			transactionDetail.Quantity = uint(cartItem.Quantity)
			transactionDetail.PackagePrice = packagePriceMap[cartItem.ProductPackageID]

			transactionDetailCreation := tx.Create(&transactionDetail)

			if transactionDetailCreation.Error != nil {
				if strings.HasPrefix(transactionDetailCreation.Error.Error(), "Error 1452") {
					return errors.New(fmt.Sprintf("no invoice id '%v' found in the transaction table", transactionDetail.InvoiceID))
				}
				return transactionDetailCreation.Error
			}
		}

		config.Db.Table("payment_methods").Where("id = ?", payment.ID).Find(&paymentMethod)
		transactionAPI.InvoiceID = invoiceId
		transactionAPI.Total = total
		transactionAPI.PaymentMethodID = payment.ID
		transactionAPI.Description = paymentMethod.Description

		return nil
	})

	if err != nil {
		return transactionAPI, err
	}

	return transactionAPI, nil
}