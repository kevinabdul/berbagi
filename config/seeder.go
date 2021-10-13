package config

import (
	"berbagi/models"
	"berbagi/utils/password"
	//datavalidation "berbagi/utils/registration"
	"errors"
	"os"
	"fmt"
	"time"
	"strconv"

	"gorm.io/gorm"

)

func InsertProvince() {
	provinces := []models.Province{{Name: "DKI Jakarta"}, {Name: "Denpasar"}}
	
	Db.Create(&provinces)
}

func InsertCity() {
	cities := []models.City{{Name: "Jakarta Pusat", ProvinceID: 1}, {Name: "Bali", ProvinceID: 2}}

	Db.Create(&cities)
}

func InsertCategory() {
	categories := []models.Category{{Name: "school uniform"}, {Name: "food"}, {Name: "school utility"}, {Name: "communication package"}}

	Db.Create(&categories)
}

func InsertProduct() {
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

func InsertProductPackage() {
	productPackage := []models.ProductPackage{{Name: "School Package-SD_Telkomsel"}, {Name: "School Package-SMP-Telkomsel"},
		{Name: "Food Package-Telur"}, {Name: "Food package-Ayam"}}

	Db.Create(&productPackage)
}

func InsertProductPackageDetail() {
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

func InsertPaymentMethod() {
	paymentMethod := []models.PaymentMethod{{Name:"alfamart", Description: "steps to pay via alfamart"},
	{Name: "bank transfer", Description:"steps to pay via bank transfer"}, {Name: "gopay", Description: "steps to pay via gopay"}}

	Db.Create(&paymentMethod)
}

func InsertAction() {
	actions := []models.Action{{Name:"GET"}, {Name:"POST"}, {Name:"PUT"}, {Name:"DELETE"}}
	Db.Create(&actions)
}

func InsertResource() {
	resources := []models.Resource{{Path:"/product-carts"}, {Path:"/checkout"}, {Path:"/payments"}, {Path:"/gifts"},
	{Path:"/location"}, {Path:"/request/donations"}, {Path:"/request/gifts"}, {Path:"/request/services"},
	{Path:"/proficiencies"}, {Path:"/volunteer"}, {Path:"/services"}, {Path:"/services/verification"}}
	Db.Create(&resources)
}

func InsertPermission() {
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

func InsertRole() {
	roles := []models.Role{{Name:"admin"}, {Name:"donor"}, {Name:"volunteer"}, {Name:"children"}, {Name:"foundation"}}
	Db.Create(&roles)
}

func InsertRolePermission() {
	rolePermissions := []models.RolePermission{{RoleID:2, PermissionID:1}, {RoleID:2, PermissionID:2}, {RoleID:2, PermissionID:3},
	{RoleID:2, PermissionID:4}, {RoleID:2, PermissionID:5}, {RoleID:2, PermissionID:6}, {RoleID:4, PermissionID:7},
	{RoleID:2, PermissionID:8}, {RoleID:3, PermissionID:8}, {RoleID:5, PermissionID:9}, {RoleID:4, PermissionID:10},
	{RoleID:5, PermissionID:11}, {RoleID:3, PermissionID:12}, {RoleID:5, PermissionID:12}, {RoleID:3, PermissionID:13},
	{RoleID:3, PermissionID:14}, {RoleID:3, PermissionID:15}, {RoleID:3, PermissionID:16}, {RoleID:3, PermissionID:17}}
	Db.Create(&rolePermissions)
}

type RegistrationAPI struct {
	UserID		 	uint   `json:"user_id"`
	Name         	string `json:"name"`
	Email        	string `json:"email"`
	Password     	string `json:"password"`
	NIK          	string `json:"nik"`
	BirthDate	 	string `json:"birth_date"`
	AddressName  	string `json:"address_name"`
	Latitude     	string `json:"latitude"`
	Longitude    	string `json:"longitude"`
	CityID       	uint   `json:"city_id"`
	ProvinceID   	uint   `json:"province_id"`
	ProficiencyID	uint   `json:"proficiency_id"`
	LicenseID    	uint   `json:"license_id"`
	RoleID         	uint `json:"role_id"`
	AdminKey 		string `json:"admin_key"`
}

/* test seeder function */

func InsertUser(incomingData models.RegistrationAPI) (int,error) {
	hashedPassword,_ := password.Hash(incomingData.Password)
	newUser := models.User{}

	transactionErr := Db.Transaction(func(tx *gorm.DB) error {

		newAddress := models.Address{}
		newAddress.Name = incomingData.AddressName
		newAddress.Latitude = incomingData.Latitude
		newAddress.Longitude = incomingData.Longitude
		newAddress.CityID = incomingData.CityID
		newAddress.ProvinceID = incomingData.ProvinceID

		if err := tx.Model(models.Address{}).Create(&newAddress).Error; err != nil {
			return err
		}

		newUser.Name = incomingData.Name
		newUser.NIK = incomingData.NIK
		newUser.Email = incomingData.Email
		newUser.Password = hashedPassword
		newUser.RoleID = incomingData.RoleID

		if err := tx.Model(models.User{}).Create(&newUser).Error; err != nil {
			return err
		}

		if incomingData.RoleID == 2 {
			newUserRole := models.Donor{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("donors").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 1 {
			adminKey := os.Getenv("ADMIN_KEY")

			if adminKey != incomingData.AdminKey || incomingData.AdminKey == "" {
				return errors.New("Invalid admin key")
			}

			newUserRole := models.Admin{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("admins").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 4 {
			newUserRole := models.Children{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("childrens").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 3 {
			newUserRole := models.Volunteer{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.ProficiencyID = incomingData.ProficiencyID
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("volunteers").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 5 {
			newUserRole := models.Foundation{}
			newUserRole.UserID = newUser.ID
			newUserRole.LicenseID = incomingData.LicenseID
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("foundations").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		}

		return nil
	})

	if transactionErr != nil {
		return -1, transactionErr
	}

	return int(newUser.ID), nil
}

func InsertProductCart(userCart []models.ProductCart, donorId int) {
	for _, cartItem := range userCart {
		targetCart := models.ProductCart{}
		Db.Where(models.ProductCart{DonorID: uint(donorId), RecipientID: uint(cartItem.RecipientID),
		ProductPackageID: cartItem.ProductPackageID}).Assign(models.ProductCart{Quantity: cartItem.Quantity}).FirstOrCreate(&targetCart)
	}
}

func CheckoutProductCart(payment models.PaymentMethod,donorId int) models.Transaction{
	var productCart []models.GiftAPI
	Db.Table("product_carts").
	Select("product_carts.recipient_id, product_carts.product_package_id, product_carts.quantity").
	Where(`product_carts.donor_id = ?`, donorId).Find(&productCart)

	Db.Table("product_carts").Where("donor_id = ?", donorId).Unscoped().Delete(&models.ProductCart{})

	packageQuantity := map[uint]int{}
	packagetarget := ""
	dictRecipient := map[uint][]models.GiftAPI{}

	for _, v := range productCart {
		if _, ok := packageQuantity[v.ProductPackageID]; !ok {
			packageQuantity[v.ProductPackageID] = v.Quantity
		} else {
			packageQuantity[v.ProductPackageID] += v.Quantity
		}

		if _, ok := dictRecipient[v.RecipientID]; !ok {
			dictRecipient[v.RecipientID] = append(dictRecipient[v.RecipientID], v)
		}
	}

	for k, _ := range packageQuantity {
		packagetarget += strconv.Itoa(int(k)) + ","
	}

	packagetarget = "(" + packagetarget[0: len(packagetarget) - 1] + ")"

	packagePrices := []models.PackagePrice{}
	packagePriceMap := map[uint]int{}

	Db.Table("product_package_details").
	Select("product_package_details.product_package_id, sum(products.price) as price").
	Joins("left join products on products.id = product_package_details.product_id").Group("product_package_details.product_package_id").
	Having(fmt.Sprintf("product_package_id in %s", packagetarget)).Find(&packagePrices)

	var total int

	for _, v := range packagePrices {
		total += v.Price * int(packageQuantity[v.ProductPackageID])
		packagePriceMap[v.ProductPackageID] = v.Price
	}

	invoiceId := fmt.Sprintf("BERBAGI.DONOR.%03v.%v", donorId, time.Now().String()[0:19])

	transaction := models.Transaction{}
	transaction.DonorID = uint(donorId)
	transaction.InvoiceID = invoiceId
	transaction.PaymentMethodID = payment.ID
	transaction.Total = total

	Db.Create(&transaction)

	transactionDetail := models.TransactionDetail{}

	transactionDetail.InvoiceID = invoiceId

	for _, cartItem := range productCart {
		transactionDetail.RecipientID = cartItem.RecipientID
		transactionDetail.ProductPackageID = cartItem.ProductPackageID
		transactionDetail.Quantity = uint(cartItem.Quantity)
		transactionDetail.PackagePrice = packagePriceMap[cartItem.ProductPackageID]

		Db.Create(&transactionDetail)
	}

	return transaction
}

func ResolveOnePayment(transaction models.Transaction) {
	Db.Table("transactions").Where("invoice_id = ?", transaction.InvoiceID).Updates(models.Transaction{PaymentStatus:"paid"})
}
