package models

type CheckoutGetResponse struct {
	Recipients	 		[]RecipientGift 	`json:"recipients"`
	PackageList 		[]PackageListAPI 	`json:"package_lists"`
	PaymentMethods 		[]PaymentMethod		`json:"payment_methods"`
}