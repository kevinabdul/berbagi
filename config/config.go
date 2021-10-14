package config

import (
	"berbagi/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db       *gorm.DB
	dbTables = map[string]interface{}{
		"provinces":                &models.Province{},
		"cities":                   &models.City{},
		"addresses":                &models.Address{},
		"proficiencies":            &models.Proficiency{},
		"users":                    &models.User{},
		"donors":                   &models.Donor{},
		"categories":               &models.Category{},
		"products":                 &models.Product{},
		"product_carts":            &models.ProductCart{},
		"product_packages":         &models.ProductPackage{},
		"product_package_details":  &models.ProductPackageDetail{},
		"payment_methods":          &models.PaymentMethod{},
		"volunteers":               &models.Volunteer{},
		"childrens":                &models.Children{},
		"foundations":              &models.Foundation{},
		"admins":                   &models.Admin{},
		"service_carts":            &models.ServiceCart{},
		"confirm_services":         &models.ConfirmService{},
		"transactions":             &models.Transaction{},
		"transaction_details":      &models.TransactionDetail{},
		"completions":              &models.Completion{},
		"certificates":             &models.Certificate{},
		"actions":                  &models.Action{},
		"resources":                &models.Resource{},
		"roles":                    &models.Role{},
		"permissions":              &models.Permission{},
		"role_permissions":         &models.RolePermission{},
		"requests":                 &models.Request{},
		"gift_request_details":     &models.GiftRequestDetails{},
		"donation_request_details": &models.DonationRequestDetails{},
		"service_request_details":  &models.ServiceRequestDetails{},
		"donations":                &models.Donation{},
		"donation_carts":           &models.DonationCart{}}
)

func InitDb() {
	err1 := godotenv.Load("./.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err2 error
	Db, err2 = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}

	Db.Migrator().DropTable("provinces")
	Db.Migrator().DropTable("cities")
	Db.Migrator().DropTable("addresses")
	Db.Migrator().DropTable("donors")
	Db.Migrator().DropTable("categories")
	Db.Migrator().DropTable("products")
	Db.Migrator().DropTable("product_carts")
	Db.Migrator().DropTable("product_packages")
	Db.Migrator().DropTable("product_package_details")
	Db.Migrator().DropTable("payment_methods")
	Db.Migrator().DropTable("volunteers")
	Db.Migrator().DropTable("proficiencies")
	Db.Migrator().DropTable("childrens")
	Db.Migrator().DropTable("foundations")
	Db.Migrator().DropTable("users")
	Db.Migrator().DropTable("admins")
	Db.Migrator().DropTable("transaction_details")
	Db.Migrator().DropTable("transactions")
	Db.Migrator().DropTable("actions")
	Db.Migrator().DropTable("resources")
	Db.Migrator().DropTable("permissions")
	Db.Migrator().DropTable("roles")
	Db.Migrator().DropTable("role_permissions")
	Db.Migrator().DropTable(&models.Certificate{})
	Db.Migrator().DropTable(&models.Completion{})
	Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
	Db.Migrator().DropTable(&models.ServiceCart{})
	Db.Migrator().DropTable(&models.Request{})
	Db.Migrator().DropTable(&models.GiftRequestDetails{})
	Db.Migrator().DropTable(&models.Donation{})
	Db.Migrator().DropTable(&models.DonationCart{})
	Db.Migrator().DropTable(&models.DonationRequestDetails{})
	Db.Migrator().DropTable(&models.ServiceRequestDetails{})
	Db.Migrator().DropTable(&models.TransactionDonationDetail{})
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Province{})
	Db.AutoMigrate(&models.City{})
	Db.AutoMigrate(&models.Address{})
	Db.AutoMigrate(&models.Proficiency{})
	Db.AutoMigrate(&models.Admin{})
	Db.AutoMigrate(&models.Donor{})
	Db.AutoMigrate(&models.Category{})
	Db.AutoMigrate(&models.Product{})
	Db.AutoMigrate(&models.ProductCart{})
	Db.AutoMigrate(&models.ProductPackage{})
	Db.AutoMigrate(&models.ProductPackageDetail{})
	Db.AutoMigrate(&models.PaymentMethod{})
	Db.AutoMigrate(&models.Volunteer{})
	Db.AutoMigrate(&models.Children{})
	Db.AutoMigrate(&models.Foundation{})
	Db.AutoMigrate(&models.ServiceCart{})
	Db.AutoMigrate(&models.ConfirmServicesAPI{})
	Db.AutoMigrate(&models.Transaction{})
	Db.AutoMigrate(&models.TransactionDetail{})
	Db.AutoMigrate(&models.Completion{})
	Db.AutoMigrate(&models.Certificate{})
	Db.AutoMigrate(&models.Action{})
	Db.AutoMigrate(&models.Resource{})
	Db.AutoMigrate(&models.Role{})
	Db.AutoMigrate(&models.Permission{})
	Db.AutoMigrate(&models.RolePermission{})
	Db.AutoMigrate(&models.Request{})
	Db.AutoMigrate(&models.GiftRequestDetails{})
	Db.AutoMigrate(&models.DonationRequestDetails{})
	Db.AutoMigrate(&models.ServiceRequestDetails{})
	Db.AutoMigrate(&models.Donation{})
	Db.AutoMigrate(&models.DonationCart{})
	Db.AutoMigrate(&models.TransactionDonationDetail{})

	InsertProvince()

	InsertCity()

	InsertCategory()

	InsertProduct()

	InsertProductPackage()

	InsertProductPackageDetail()

	InsertPaymentMethod()

	InsertAction()

	InsertResource()

	InsertPermission()

	InsertRole()

	InsertRolePermission()
}

// this config for API testing purpose
func InitDBTest(tables ...string) {
	if err := godotenv.Load("./../.env"); err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file. Got this error: %v", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME_TEST"))

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest(tables...)
}

func InitMigrateTest(tables ...string) {
	for _, v := range tables {
		Db.Migrator().DropTable(dbTables[v])
	}
	for _, v := range tables {
		Db.AutoMigrate(dbTables[v])
	}
}
