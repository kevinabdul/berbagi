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
	Db *gorm.DB
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
	Db.Migrator().DropTable("proficiencies")
	Db.Migrator().DropTable("donors")
	Db.Migrator().DropTable("categories")
	Db.Migrator().DropTable("products")
	Db.Migrator().DropTable("product_carts")
	Db.Migrator().DropTable("product_packages")
	Db.Migrator().DropTable("product_package_details")
	Db.Migrator().DropTable("payment_methods")
	Db.Migrator().DropTable("volunteers")
	Db.Migrator().DropTable("childrens")
	Db.Migrator().DropTable("foundations")
	Db.Migrator().DropTable("users")
	Db.Migrator().DropTable("admins")
	Db.Migrator().DropTable("service_carts")
	Db.Migrator().DropTable("confirm_services")
	Db.Migrator().DropTable("admins")
	Db.Migrator().DropTable("transaction_details")
	Db.Migrator().DropTable("transactions")
	Db.AutoMigrate(&models.Province{})
	Db.AutoMigrate(&models.City{})
	Db.AutoMigrate(&models.Address{})
	Db.AutoMigrate(&models.Proficiency{})
	Db.AutoMigrate(&models.User{})
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
	Db.AutoMigrate(&models.Admin{})
	Db.AutoMigrate(&models.ServiceCart{})
	Db.AutoMigrate(&models.ConfirmService{})
	Db.AutoMigrate(&models.Transaction{})
	Db.AutoMigrate(&models.TransactionDetail{})
  
	insertProvince()

	insertCity()

	insertCategory()

	insertProduct()

	insertProductPackage()

	insertProductPackageDetail()

	insertPaymentMethod()
}

