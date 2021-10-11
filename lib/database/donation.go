package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func MakeDonationToCart(donation models.DonationInputData) (models.NewDonationResponseAPI, error) {
	var cart models.DonationCart
	var res models.NewDonationResponseAPI

	cart.DonorID = donation.DonorID
	cart.RecipientID = donation.RecipientID
	cart.Amount = donation.Amount

	whereQuery := fmt.Sprintf("donor_id = %d AND recipient_id = %d",
					cart.DonorID, cart.RecipientID)

	if donation.RequestID > 0 {
		cart.RequestID = donation.RequestID
		res.RequestID = donation.RequestID
		whereQuery += fmt.Sprintf(" AND request_id = %d", cart.RequestID)
	}

	if config.Db.Model(&cart).Where(whereQuery).Updates(&cart).RowsAffected == 0 {
		tx := config.Db.Create(&cart)
		if tx.Error != nil {
			return models.NewDonationResponseAPI{}, tx.Error
		}
	}

	res.DonorID = cart.DonorID
	res.RecipientID = cart.RecipientID
	res.Amount = cart.Amount

	return res, nil
}

func GetUnresolvedDonations(userId uint) ([]models.DonationResponse, error) {
	var donation []models.DonationResponse
	queryField := `donation_id, donor_id, recipient_id, 
				request_id, amount, updated_at as made_at`

	tx := config.Db.Table("donations").Select(queryField).Where(
		"payment_status = false").Scan(&donation)
	if tx.Error != nil {
		return []models.DonationResponse{}, tx.Error
	}
	return donation, nil
}

func GetResolvedDonations(userId uint) ([]models.DonationResponse, error) {
	var donation []models.DonationResponse
	queryField := `donation_id, donor_id, recipient_id, 
				request_id, amount, updated_at as resolved_at`

	tx := config.Db.Table("donations").Select(queryField).Where(
		"payment_status = true").Scan(&donation)
	if tx.Error != nil {
		return []models.DonationResponse{}, tx.Error
	}
	return donation, nil
}

func ListDonationInCart(userId uint) ([]models.DonationResponse, error) {
	var donation []models.DonationResponse
	queryField := `0 as donation_id, donor_id, recipient_id, 
				request_id, amount, updated_at as added_at`

	tx := config.Db.Table("donation_carts").Select(queryField).Scan(&donation)
	if tx.Error != nil {
		return []models.DonationResponse{}, tx.Error
	}
	return donation, nil
}

func DeleteDonationFromCart(data models.DonationInputData) (int, error) {
	tx := config.Db.Where(`donor_id = ? AND
			recipient_id = ? AND
			request_id = ?`,
		data.DonorID,
		data.RecipientID,
		data.RequestID).Unscoped().Delete(&models.DonationCart{})

	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, nil
	}
	return int(tx.RowsAffected), nil
}

func UpdateDonationInCart(data models.DonationInputData) error {
	tx := config.Db.Model(&models.DonationCart{}).Where(
		`donor_id = ? AND
		recipient_id = ? AND
		request_id = ?`,
		data.DonorID,
		data.RecipientID,
		data.RequestID).Update("amount", data.Amount)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetDonationInCart(data models.DonationInputData) (models.DonationResponse, error) {
	var cart models.DonationCart

	query := fmt.Sprintf("donor_id = %d AND recipient_id = %d",
			 data.DonorID, data.RecipientID)
	tx := config.Db.Model(cart).Where(query).Find(&cart)

	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	res := models.DonationResponse{
		// DonationID: 0,
		DonorID:     cart.DonorID,
		RecipientID: cart.RecipientID,
		RequestID:   cart.RequestID,
		Amount:      cart.Amount,
		MadeAt:      cart.UpdatedAt,
	}
fmt.Println("res", tx.RowsAffected, res)
	return res, nil
}

// func CheckoutDonation(data models.DonationInputData) (models.DonationResponse, error) {
// 	dataSearch := models.DonationInputData{
// 		DonorID:     data.DonorID,
// 		RecipientID: data.RecipientID,
// 		RequestID:   data.RequestID,
// 	}

// 	donation, err := GetDonationInCart(dataSearch)
// 	if err != nil {
// 		return models.DonationResponse{}, err
// 	}

// 	checkout := models.Donation{
// 		DonorID:       donation.DonorID,
// 		RecipientID:   donation.RecipientID,
// 		RequestID:     donation.RequestID,
// 		Amount:        donation.Amount,
// 		PaymentStatus: "false",
// 	}

// 	if donation.Amount == 0 {
// 		checkout.Amount = data.Amount
// 	}

// 	tx := config.Db.Create(&checkout)
// 	if tx.Error != nil {
// 		return models.DonationResponse{}, tx.Error
// 	}

// 	donation.DonationID = checkout.ID
// 	donation.PaymentStatus = checkout.PaymentStatus
// 	donation.MadeAt = checkout.UpdatedAt
// 	return donation, nil
// }

func CheckoutDonation(data models.DonationInputData, quick string) (models.TransactionAPI, error) {
	var transactionAPI models.TransactionAPI
	var paymentMethod models.PaymentMethod

	err := config.Db.Transaction(func(tx *gorm.DB) error {
		if quick != "yes" && quick != "true" && quick != "quick" {
			var donationCart models.DonationCart
			var deletedCart models.DonationCart
			findQuery := fmt.Sprintf("donor_id = %d AND recipient_id = %d", data.DonorID,
						 data.RecipientID)
			if data.RequestID > 0 {
				deletedCart.RequestID = data.RequestID
				findQuery += fmt.Sprintf(" AND request_id = %d", data.RequestID)
			}

			findCart := tx.Table("donation_carts").Where(findQuery).Find(&donationCart)
			if findCart.Error != nil {
				return errors.New("failed to find donation in cart")
			}
			if findCart.RowsAffected == 0 {
				return errors.New("add donation to cart first")
			}

			data.Amount = donationCart.Amount

			deletedCart.DonorID = data.DonorID
			deletedCart.RecipientID = data.RecipientID
	
			deleteCart := tx.Table("donation_carts").Unscoped().Delete(&deletedCart)
			if deleteCart.Error != nil {
				return deleteCart.Error
			}
			if deleteCart.RowsAffected == 0 {
				return errors.New("failed to delete donation from cart")
			}
		}

		donation := models.Donation{
			DonorID:       data.DonorID,
			RecipientID:   data.RecipientID,
			RequestID:     data.RequestID,
			Amount:        data.Amount,
			PaymentStatus: "false",
		}
		
		donationDetail := tx.Create(&donation)
		if donationDetail.Error != nil {
			return donationDetail.Error
		}

		
		invoiceId := fmt.Sprintf("BERBAGI.DONOR.%03v.DONATE.%03v.%v", data.DonorID, donation.ID, time.Now().String()[0:19])

		var transactionDetail models.TransactionDonationDetail
		transactionDetail.InvoiceID = invoiceId
		transactionDetail.DonationID = donation.ID
		transactionDetail.RecipientID = data.RecipientID
		transactionDetail.Amount = data.Amount

		transactionDetailCreation := tx.Create(&transactionDetail)
		if transactionDetailCreation.Error != nil {
			if strings.HasPrefix(transactionDetailCreation.Error.Error(), "Error 1452") {
				return errors.New(fmt.Sprintf("no invoice id '%v' found in the transaction table", transactionDetail.InvoiceID))
			}
			return transactionDetailCreation.Error
		}
		
		var transaction models.Transaction
		transaction.DonorID = data.DonorID
		transaction.InvoiceID = invoiceId
		transaction.PaymentMethodID = uint(data.PaymentID)
		transaction.Total = int(data.Amount)

		transactionCreation := tx.Create(&transaction)
		if transactionCreation.Error != nil {
			if strings.HasPrefix(transactionCreation.Error.Error(), "Error 1452") {
				return errors.New(fmt.Sprintf("no payment_method_id '%v' found in the payment method table", data.PaymentID))
			}
			return transactionCreation.Error
		}
		

		config.Db.Table("payment_methods").Where("id = ?", data.PaymentID).Find(&paymentMethod)
		transactionAPI.InvoiceID = invoiceId
		transactionAPI.Total = data.Amount
		transactionAPI.PaymentMethodID = uint(data.PaymentID)
		transactionAPI.Description = paymentMethod.Description

		return nil
	})

	if err != nil {
		return transactionAPI, err
	}

	return transactionAPI, nil
}

func GetSpecificDonation(donationId uint) (models.DonationResponse, error) {
	var donation models.Donation
	donation.ID = donationId
	tx := config.Db.Find(&donation)
	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	res := models.DonationResponse{
		DonationID:    donation.ID,
		DonorID:       donation.DonorID,
		RecipientID:   donation.RecipientID,
		RequestID:     donation.RequestID,
		Amount:        donation.Amount,
		PaymentStatus: donation.PaymentStatus,
		MadeAt:        donation.UpdatedAt,
	}
	return res, nil
}

func GetBulkDonations(userId uint, resolved string) ([]models.Donation, error) {
	var donation []models.Donation

	if resolved == "yes" {
		tx := config.Db.Where("donor_id = ? AND payment_status = true", userId).Find(&donation)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if resolved == "no" {
		tx := config.Db.Where("donor_id = ? AND payment_status = false", userId).Find(&donation)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := config.Db.Where("donor_id = ?", userId).Find(&donation)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return donation, nil
}

func ChangePaymentStatusToPaid(donationId uint, paid string) (models.DonationResponse, error) {
	var donation models.Donation

	donation.ID = donationId
	tx := config.Db.Find(&donation).Update("payment_status", paid)
	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	res := models.DonationResponse{
		DonationID:    donation.ID,
		DonorID:       donation.DonorID,
		RecipientID:   donation.RecipientID,
		RequestID:     donation.RequestID,
		Amount:        donation.Amount,
		PaymentStatus: donation.PaymentStatus,
		MadeAt:        donation.UpdatedAt,
	}
	return res, nil
}
