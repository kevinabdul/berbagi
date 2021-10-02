package config

import (
	"berbagi/models"
)

func insertProvince() {
	provinces := []models.Province{{Name: "DKI Jakarta"}, {Name: "Denpasar"}}
	
	Db.Create(&provinces)
}

func insertCity() {
	cities := []models.City{{Name: "Jakarta Pusat", ProvinceID: 1}, {Name: "Bali", ProvinceID: 2}}

	Db.Create(&cities)
}

func insertCategory() {
	categories := []models.Category{{Name: "school uniform"}, {Name: "food"}, {Name: "school utility"}, {Name: "communication package"}}

	Db.Create(&categories)
}

func insertProduct() {
	products := []models.Product{{Name: "Seragam SD", Price: 60000, CategoryID: 1}, {Name: "Seragam SMP", Price: 75000, CategoryID: 1},
		{Name: "Beras 5Kg", Price: 65000, CategoryID: 2}, {Name: "Telur Ayam (10)", Price: 13000, CategoryID: 2},
		{Name: "Daging Ayam 1Kg", Price: 35000, CategoryID: 2}, {Name: "Minyak Sayur 2L", Price: 24000, CategoryID: 2},
		{Name: "Susu Kotak 1L", Price: 15000, CategoryID: 2}, {Name: "Gula 1Kg", Price: 10000, CategoryID: 2},
		{Name: "Buku Paket SD", Price: 200000, CategoryID: 3}, {Name: "Buku Paket SMP", Price: 300000, CategoryID: 3},
		{Name: "Sepatu Sekolah", Price: 180000, CategoryID: 3}, {Name: "Tas Sekolah", Price: 135000, CategoryID: 3},
		{Name: "Kuota Data 20GB Telkomsel", Price: 100000, CategoryID: 4}, {Name: "Kuota data 50GB Indosat", Price: 100000, CategoryID: 4},
		{Name: "Pulsa 100000 Telkomsel", Price: 100000, CategoryID: 4}, {Name: "Pulsa 100000 Indosat", Price: 100000, CategoryID: 4}}

	Db.Create(&products)
}

func insertProductPackage() {
	productPackage := []models.ProductPackage{{Name: "School Package-SD_Telkomsel"}, {Name: "School Package-SMP-Telkomsel"},
		{Name: "Food Package-Telur"}, {Name: "Food package-Ayam"}}

	Db.Create(&productPackage)
}

func insertProductPackageDetail() {
	productPackageDetail := []models.ProductPackageDetail{{ProductPackageID: 1, ProductID: 1, Quantity: 1},
		{ProductPackageID: 1, ProductID: 9, Quantity: 1}, {ProductPackageID: 1, ProductID: 11, Quantity: 1},
		{ProductPackageID: 2, ProductID: 2, Quantity: 1}, {ProductPackageID: 2, ProductID: 9, Quantity: 1},
		{ProductPackageID: 2, ProductID: 11, Quantity: 1}, {ProductPackageID: 3, ProductID: 3, Quantity: 1},
		{ProductPackageID: 3, ProductID: 4, Quantity: 1}, {ProductPackageID: 3, ProductID: 6, Quantity: 1},
		{ProductPackageID: 3, ProductID: 7, Quantity: 1}, {ProductPackageID: 3, ProductID: 8, Quantity: 1},
		{ProductPackageID: 4, ProductID: 3, Quantity: 1}, {ProductPackageID: 4, ProductID: 5, Quantity: 1},
		{ProductPackageID: 4, ProductID: 6, Quantity: 1}, {ProductPackageID: 4, ProductID: 7, Quantity: 1},
		{ProductPackageID: 4, ProductID: 8, Quantity: 1}}

	Db.Create(&productPackageDetail)
}

func insertPaymentMethod() {
	paymentMethod := []models.PaymentMethod{{Name:"alfamart", Description: "steps to pay via alfamart"},
	{Name: "bank transfer", Description:"steps to pay via bank transfer"}, {Name: "gopay", Description: "steps to pay via gopay"}}

	Db.Create(&paymentMethod)
}