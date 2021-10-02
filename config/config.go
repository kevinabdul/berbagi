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
	Db.Migrator().DropTable("volunteers")
	Db.Migrator().DropTable("childrens")
	Db.Migrator().DropTable("foundations")
	Db.Migrator().DropTable("users")
	Db.Migrator().DropTable("admins")
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
	Db.AutoMigrate(&models.Volunteer{})
	Db.AutoMigrate(&models.Children{})
	Db.AutoMigrate(&models.Foundation{})
	Db.AutoMigrate(&models.Admin{})
	Db.AutoMigrate(&models.ServiceCart{})
	Db.AutoMigrate(&models.ConfirmServicesAPI{})
  
	insertProvince()

	insertCity()

	insertCategory()

	insertProduct()

	insertProductPackage()

	insertProductPackageDetail()
}

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
	products := []models.Product{{Name: "Seragam SD", Price: 70000, CategoryID: 1}, {Name: "Seragam SMP", Price: 70000, CategoryID: 1},
		{Name: "Beras 5Kg", Price: 65000, CategoryID: 2}, {Name: "Telur Ayam (10)", Price: 13000, CategoryID: 2},
		{Name: "Daging Ayam 1Kg", Price: 35000, CategoryID: 2}, {Name: "Minyak Sayur 2L", Price: 24000, CategoryID: 2},
		{Name: "Susu Kotak 1L", Price: 15000, CategoryID: 2}, {Name: "Gula 1Kg", Price: 10000, CategoryID: 2},
		{Name: "Kuota Data 20GB Telkomsel", Price: 100000, CategoryID: 3}, {Name: "Kuota data 50GB Indosat", Price: 100000, CategoryID: 3},
		{Name: "Pulsa 100000 Telkomsel", Price: 100000, CategoryID: 3}, {Name: "Pulsa 100000 Indosat", Price: 100000, CategoryID: 3}}
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
