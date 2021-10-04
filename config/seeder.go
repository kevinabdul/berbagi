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

func insertAction() {
	actions := []models.Action{{Name:"GET"}, {Name:"POST"}, {Name:"PUT"}, {Name:"DELETE"}}
	Db.Create(&actions)
}

func insertResource() {
	resources := []models.Resource{{Path:"/product-carts"}, {Path:"/checkout"}, {Path:"/payments"}, {Path:"/gifts"},
	{Path:"/location"}, {Path:"/request/donations"}, {Path:"/request/gifts"}, {Path:"/request/services"},
	{Path:"/proficiencies"}, {Path:"/volunteer"}, {Path:"/services"}, {Path:"/services/verification"}}
	Db.Create(&resources)
}

func insertPermission() {
	permissions := []models.Permission{
		{ActionID:1, ResourceID:1},  // 1.  GET 	product-carts
		{ActionID:3, ResourceID:1},  // 2.  PUT 	product-carts
		{ActionID:1, ResourceID:2},	 // 3.  GET 	checkout
		{ActionID:2, ResourceID:2},  // 4.  POST 	checkout
		{ActionID:1, ResourceID:3},  // 5.  GET 	payments
		{ActionID:2, ResourceID:3},  // 6.  POST 	payments
		{ActionID:1, ResourceID:4},  // 7.  GET  	gifts
		{ActionID:1, ResourceID:5},  // 8.  GET  	location
		{ActionID:2, ResourceID:6},  // 9.  POST 	request/donations
		{ActionID:2, ResourceID:7},  // 10. POST 	request/gifts
		{ActionID:2, ResourceID:8},  // 11. POST 	request/services
		{ActionID:1, ResourceID:9},  // 12. GET  	proficiencies
		{ActionID:1, ResourceID:10}, // 13. GET  	volunteer (profile)
		{ActionID:1, ResourceID:11}, // 14. GET  	services
		{ActionID:3, ResourceID:11}, // 15. PUT  	services
		{ActionID:4, ResourceID:11}, // 16. DELETE  services
		{ActionID:1, ResourceID:12}} // 17. GET  	services/verification
	Db.Create(&permissions)
}

func insertRole() {
	roles := []models.Role{{Name:"admin"}, {Name:"donor"}, {Name:"volunteer"}, {Name:"children"}, {Name:"foundation"}}
	Db.Create(&roles)
}

func insertRolePermission() {
	rolePermissions := []models.RolePermission{{RoleID:2, PermissionID:1}, {RoleID:2, PermissionID:2}, {RoleID:2, PermissionID:3},
	{RoleID:2, PermissionID:4}, {RoleID:2, PermissionID:5}, {RoleID:2, PermissionID:6}, {RoleID:4, PermissionID:7},
	{RoleID:2, PermissionID:8}, {RoleID:3, PermissionID:8}, {RoleID:5, PermissionID:9}, {RoleID:4, PermissionID:10},
	{RoleID:5, PermissionID:11}, {RoleID:3, PermissionID:12}, {RoleID:5, PermissionID:12}, {RoleID:3, PermissionID:13},
	{RoleID:3, PermissionID:14}, {RoleID:3, PermissionID:15}, {RoleID:3, PermissionID:16}, {RoleID:3, PermissionID:17}}
	Db.Create(&rolePermissions)
}

