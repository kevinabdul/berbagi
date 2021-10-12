package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
	"fmt"
	//"strings"
	"strconv"
)

func GetPendingPaymentsByDonorId(donorId int) ([]models.PendingPaymentAPI, error){
	invoiceMap := map[string]map[string][]models.TransactionDetailAPI{}

	paymentDetails := []models.TransactionDetailAPI{}

	pendingPaymentSearchRes := config.Db.Table("transactions").
	Select(`transactions.invoice_id, transaction_details.product_package_id, transaction_details.recipient_id, transactions.donor_id,
	transactions.payment_status, payment_methods.name as payment_method_name, transaction_details.package_price * transaction_details.quantity as total,
	transactions.payment_method_id, payment_methods.description, transaction_details.package_price, transaction_details.quantity`).
	Joins("left join transaction_details on transactions.invoice_id = transaction_details.invoice_id").
	Joins("left join payment_methods on transactions.payment_method_id = payment_methods.id").
	Where("donor_id = ? and payment_status = ?", donorId, "pending").Find(&paymentDetails)

	if pendingPaymentSearchRes.Error != nil {
		return []models.PendingPaymentAPI{}, pendingPaymentSearchRes.Error
	}

	if pendingPaymentSearchRes.RowsAffected == 0 {
		return []models.PendingPaymentAPI{}, errors.New("No pending payments found")
	}

	for _, paymentDetail := range paymentDetails {
		if _, ok := invoiceMap[paymentDetail.InvoiceID]; !ok {
			newDetail := map[string][]models.TransactionDetailAPI{"detail": []models.TransactionDetailAPI{}}
			invoiceMap[paymentDetail.InvoiceID] = newDetail
			invoiceMap[paymentDetail.InvoiceID]["detail"] = append(invoiceMap[paymentDetail.InvoiceID]["detail"], paymentDetail)
		} else {
			invoiceMap[paymentDetail.InvoiceID]["detail"] = append(invoiceMap[paymentDetail.InvoiceID]["detail"], paymentDetail)
		}
	}
	
	res := []models.PendingPaymentAPI{}

	for id, invoice := range invoiceMap {
		result := models.PendingPaymentAPI{}
		result.InvoiceID = id
		result.PaymentMethodID =  invoice["detail"][0].PaymentMethodID
		result.PaymentMethodName = invoice["detail"][0].PaymentMethodName
		result.Description = invoice["detail"][0].Description
		result.Detail = invoice["detail"]

		for _, detail := range invoice["detail"] {
			result.Total += detail.Total
		}

		res = append(res, result)
	}

	return res, nil
}

func AddPendingPaymentByDonorId(payment models.UserPaymentAPI, donorId int) (models.ReceiptAPI, error) {
	transactionTarget := []models.TransactionDetailAPI{}

	findTransaction := config.Db.Table("transactions").
	Select(`transactions.invoice_id, transactions.payment_status, 
	transaction_details.product_package_id, transaction_details.package_price * transaction_details.quantity as total, 
	transactions.payment_method_id, payment_methods.name, payment_methods.description`).
	Joins("left join transaction_details on transactions.invoice_id = transaction_details.invoice_id").
	Joins("left join payment_methods on transactions.payment_method_id = payment_methods.id").
	Where("transaction_details.invoice_id = ?", payment.InvoiceID).Find(&transactionTarget)

	if findTransaction.Error != nil {
		return models.ReceiptAPI{}, findTransaction.Error
	}

	if findTransaction.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New("No invoice_id found")
	}

	if transactionTarget[0].PaymentStatus != "pending" {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified invoice id: %v has been %v", payment.InvoiceID, 
		transactionTarget[0].PaymentStatus))	
	}

	if payment.PaymentMethodID != transactionTarget[0].PaymentMethodID {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified payment method doesnt match. Please pay using payment_method_id: %v", 
		transactionTarget[0].PaymentMethodID))
	}

	var total int

	for _, transactionDetail := range transactionTarget {
		total += transactionDetail.Total
	}

	if total != payment.Total {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified total amount doesnt match. Total mount to be paid: %v",
		total))
	}

	newTransaction := models.Transaction{}

	updateRes := config.Db.Model(&newTransaction).Where("invoice_id = ?", payment.InvoiceID).
	Updates(models.Transaction{PaymentStatus:"paid"})

	if updateRes.Error != nil {
		return models.ReceiptAPI{}, updateRes.Error
	}

	if updateRes.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Failed to update invoice: %v", payment.InvoiceID))
	}

	receipt := models.ReceiptAPI{}

	receipt.DonorID = strconv.Itoa(donorId)
	receipt.InvoiceID = payment.InvoiceID
	receipt.Total = total
	receipt.PaymentMethodID = payment.PaymentMethodID
	receipt.CreatedAt = newTransaction.UpdatedAt

	return receipt, nil
}

func AddPendingDonationPaymentByDonorId(payment models.UserPaymentAPI, donorId int) (models.ReceiptAPI, error) {
	transactionTarget := models.TransactionDonationDetailAPI{}
	// Test Query
	// whereTestQueryFind := fmt.Sprintf("transaction_donation_details.invoice_id LIKE \"%s%%\"", payment.InvoiceID)
	// whereTestQueryUpdate := fmt.Sprintf("invoice_id LIKE \"%s%%\"", payment.InvoiceID)

	findTransaction := config.Db.Table("transactions").
	Select(`transactions.invoice_id, transactions.payment_status, 
	transaction_donation_details.purpose, transaction_donation_details.amount, 
	transactions.payment_method_id, payment_methods.name, payment_methods.description`).
	Joins("left join transaction_donation_details on transactions.invoice_id = transaction_donation_details.invoice_id").
	Joins("left join payment_methods on transactions.payment_method_id = payment_methods.id").
	Where("transaction_donation_details.invoice_id = ?", payment.InvoiceID).Find(&transactionTarget)
	// Where(whereTestQueryFind).Find(&transactionTarget) // for test purpose


	if findTransaction.Error != nil {
		return models.ReceiptAPI{}, findTransaction.Error
	}

	if findTransaction.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New("No invoice_id found")
	}

	if transactionTarget.PaymentStatus != "pending" {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified invoice id: %v has been %v", payment.InvoiceID, 
		transactionTarget.PaymentStatus))	
	}

	if payment.PaymentMethodID != transactionTarget.PaymentMethodID {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified payment method doesnt match. Please pay using payment_method_id: %v", 
		transactionTarget.PaymentMethodID))
	}

	if transactionTarget.Amount != payment.Total {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified total amount doesnt match. Total amount to be paid: %v",
		transactionTarget.Amount))
	}

	newTransaction := models.Transaction{}

	// updateRes := config.Db.Model(&newTransaction).Where(whereTestQueryUpdate). // for test purpose
	updateRes := config.Db.Model(&newTransaction).Where("invoice_id = ?", payment.InvoiceID).
	Updates(models.Transaction{PaymentStatus:"paid"})

	if updateRes.Error != nil {
		return models.ReceiptAPI{}, updateRes.Error
	}

	if updateRes.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Failed to update invoice: %v", payment.InvoiceID))
	}

	receipt := models.ReceiptAPI{}

	receipt.DonorID = strconv.Itoa(donorId)
	receipt.InvoiceID = payment.InvoiceID
	receipt.Total = transactionTarget.Amount
	receipt.PaymentMethodID = payment.PaymentMethodID
	receipt.CreatedAt = newTransaction.UpdatedAt

	return receipt, nil
}