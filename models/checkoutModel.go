package models

type CheckoutGetResponse struct {
	Recipients	 		[]RecipientGift 	`json:"recipients"`
	PackageList 		[]PackageListAPI 	`json:"package_list"`
	PaymentOptions 		[]PaymentOption		`json:"payment_options"`
}