package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID        	uint           `gorm:"primaryKey" json:"id"`
	Name      	string         `gorm:"type:varchar(50)" json:"name"`
	Description string 			`gorm:"type:varchar(1000)" json:"description"`
	CreatedAt 	time.Time      `json:"-"`
	UpdatedAt 	time.Time      `json:"-"`
	DeletedAt 	gorm.DeletedAt `gorm:"index" json:"-"`
}

// Response struct to be returned for all transaction with pending payments status from a given user
// Used in GetPendingPaymentsByUserId
type PendingPaymentAPI struct {
	InvoiceID			string				`json:"invoice_id"`
	Total 				int 				`json:"total"`
	PaymentMethodID 	uint  				`json:"payment_method_id"`
	PaymentMethodName 	string  			`json:"payment_method_name"`
	Description   		string 				`json:"description"` 
	Detail  			[]TransactionDetailAPI	`json:"detail"`
}

// Placeholder of information sent by the user when they do a payment through post payments endpoint
type UserPaymentAPI struct {
	InvoiceID			string 				`json:"invoice_id"`
	Total    			int  				`json:"total"`
	PaymentMethodID		uint 				`json:"payment_method_id"`
}


// Response struct returned to user after they completed a payment
type ReceiptAPI struct {
	DonorID				string 				`json:"donor_id"`
	InvoiceID			string 				`json:"invoice_id"`
	Total    			int  				`json:"total"`
	PaymentMethodID		uint 				`json:"payment_method_id"`
	CreatedAt 			time.Time 			`json:"payment_date"`
}
