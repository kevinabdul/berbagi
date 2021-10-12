package models

type PendingGiftAPI struct {
	DonorID				string 		`json:"donor_id"`
	RecipientID  		uint  		`json:"recipient_id"`
	InvoiceID			string		`json:"invoice_id"`
	PaymentStatus 		string		`json:"payment_status"`
	ProductPackageID	uint		`json:"product_package_id"`
	Quantity  			uint  		`json:"quantity"`
}
