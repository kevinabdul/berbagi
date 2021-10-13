package libdb

import (
	"berbagi/config"
	"berbagi/models"
	arr "berbagi/utils/arr/string"
	"errors"
	"fmt"
	//"strings"
	//"strconv"
)

func GetGiftsByChildrenId(childrenId int, paymentStatus string) ([]models.PendingGiftAPI, error){
	pendingGifts := []models.PendingGiftAPI{}

	filterCondition := fmt.Sprintf("transaction_details.recipient_id = %v", childrenId)

	if contains := arr.Contains([]string{"paid", "cancelled", "pending", "expired"}, paymentStatus); contains {
		filterCondition +=  fmt.Sprintf(" and transactions.payment_status = '%v'", paymentStatus)
	} else if !contains && paymentStatus != "" {
		return []models.PendingGiftAPI{}, errors.New("invalid payment status")	
	}

	pendingGiftSearchRes := config.Db.Table("transactions").
	Select(`transactions.invoice_id, transactions.donor_id, transaction_details.recipient_id,
	transaction_details.product_package_id, transaction_details.quantity, transactions.payment_status`).
	Joins("left join transaction_details on transactions.invoice_id = transaction_details.invoice_id").
	Where(filterCondition).Find(&pendingGifts)

	if pendingGiftSearchRes.Error != nil {
		return []models.PendingGiftAPI{}, pendingGiftSearchRes.Error
	}

	if pendingGiftSearchRes.RowsAffected == 0 {
		return []models.PendingGiftAPI{}, errors.New("no gifts found")
	}

	return pendingGifts, nil
}