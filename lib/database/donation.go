package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

func MakeDonationToCart(donation models.NewDonation) (models.NewDonationResponseAPI, error) {
	var cart models.DonationCart
	cart.DonorID = donation.DonorID
	cart.RecipientID = donation.RecipientID
	cart.RequestID = donation.RequestID
	cart.Amount = donation.Amount

	tx := config.Db.Create(&cart)
	if tx.Error != nil {
		return models.NewDonationResponseAPI{}, tx.Error
	}

	var res models.NewDonationResponseAPI
	res.DonorID = cart.DonorID
	res.RecipientID = cart.RecipientID
	res.RequestID = cart.RequestID
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

func DeleteDonationFromCart(data models.CartItemInputData) error {
	tx := config.Db.Where(`donor_id = ? AND
			recipient_id = ? AND
			request_id = ?`,
		data.DonorID,
		data.RecipientID,
		data.RequestID).Delete(&models.DonationCart{})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDonationInCart(data models.CartItemInputData) error {
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

func GetDonationInCart(data models.CartItemInputData) (models.DonationResponse, error) {
	var cart models.DonationCart

	tx := config.Db.Where(`donor_id = ? AND
		recipient_id = ? AND
		request_id = ?`,
		data.DonorID,
		data.RecipientID,
		data.RequestID).Find(&cart)

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
	return res, nil
}

func CheckoutDonation(data models.DonationCheckout) (models.DonationResponse, error) {
	dataSearch := models.CartItemInputData{
		DonorID:     data.DonorID,
		RecipientID: data.RecipientID,
		RequestID:   data.RequestID,
	}

	donation, err := GetDonationInCart(dataSearch)
	if err != nil {
		return models.DonationResponse{}, err
	}

	checkout := models.Donation{
		DonorID:       donation.DonorID,
		RecipientID:   donation.RecipientID,
		RequestID:     donation.RequestID,
		Amount:        donation.Amount,
		PaymentStatus: "false",
	}
	tx := config.Db.Create(&checkout)
	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	donation.DonationID = checkout.DonationID
	donation.PaymentStatus = checkout.PaymentStatus
	donation.MadeAt = checkout.UpdatedAt
	return donation, nil
}

func GetSpecificDonation(donationId uint) (models.DonationResponse, error) {
	var donation models.Donation
	donation.DonationID = donationId
	tx := config.Db.Find(&donation)
	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	res := models.DonationResponse{
		DonationID:    donation.DonationID,
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

	donation.DonationID = donationId
	tx := config.Db.Find(&donation).Update("payment_status", paid)
	if tx.Error != nil {
		return models.DonationResponse{}, tx.Error
	}

	res := models.DonationResponse{
		DonationID:    donation.DonationID,
		DonorID:       donation.DonorID,
		RecipientID:   donation.RecipientID,
		RequestID:     donation.RequestID,
		Amount:        donation.Amount,
		PaymentStatus: donation.PaymentStatus,
		MadeAt:        donation.UpdatedAt,
	}
	return res, nil
}
