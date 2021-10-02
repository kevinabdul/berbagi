package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	DonorID         uint        	`gorm:"primaryKey" json:"donor_id"`	
	InvoiceID		string			`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	PaymentMethodID uint           	`json:"payment_id" form:"payment_id"`
	PaymentStatus	string			`gorm:"type:enum('pending', 'unresolved', 'resolved');default:'pending'" json:"payment_status"`
	Total 			int 			`gorm:"type:int" json:"total"`
	CreatedAt 		time.Time		`json:"-"`
	UpdatedAt		time.Time		`json:"-"`
	DeletedAt       gorm.DeletedAt 	`gorm:"index" json:"-"`
	Donor 			Donor  			`gorm:"foreignKey:DonorID"`
	PaymentMethod   PaymentMethod 	`gorm:"foreignKey:PaymentMethodID"`
}

type TransactionDetail struct {
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	RecipientID  		uint  		`gorm:"primaryKey" json:"recipient_id"`
	ProductPackageID	uint		`gorm:"primaryKey;type:uint" json:"product_package_id"`
	PackagePrice  		int         `json:"package_price"`
	Quantity  			uint  		`gorm:"not null;type:smallint" json:"quantity"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
	Transaction   		Transaction 	`gorm:"foreignKey:InvoiceID;references:InvoiceID"`
	ProductPackage   	ProductPackage  `gorm:"foreignKey:ProductPackageID"`
}

// Response struct used in case of a succesful checkout in post checkout endpoint
// Succesful checkout means we are able to delete data from carts table, creating new data in transactions table,
// and moving the deleted data into transaction_details table. Any failure in those step will fail whole transaction
type TransactionAPI struct {
	InvoiceID			string		`json:"invoice_id"`
	Total 				int 		`json:"total"`
	PaymentMethodID 	uint		`json:"payment_method_id"`
	Description     	string 		`json:"description"` 		
}

// Commonly used when user try to do a payment. 
// AddPaymentByUserId will try to find corresponding transaction in a database based on UserId and information provided in UserPaymentAPI struct.
// This struct will be used as a placeholder of above query result. 
type TransactionDetailAPI struct {
	UserID				string 				`json:"user_id"`
	InvoiceID			string				`json:"invoice_id"`
	Status   			string  			`json:"status"`
	ProductName 		string				`json:"product_name"`
	ProductPrice		uint 				`json:"product_price"`
	Quantity  			uint  				`json:"quantity"`
	PaymentMethodID 	uint  				`json:"payment_method_id"`
	Description   		string 				`json:"description"` 
}
