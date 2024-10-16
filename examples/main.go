package main

import (
	"fmt"
	"github.com/pzentenoe/gorm-connection/sql"
	"log"
)

func main() {
	// Example using SQLite
	connection, err := sql.NewSQLiteConnection("test_db")
	if err != nil {
		log.Fatalf("failed to connect to SQLite database: %v", err)
	}

	// Get GORM connection
	gormConn, err := connection.GetConnection()
	if err != nil {
		log.Fatalf("failed to get GORM connection: %v", err)
	}

	// Migrate models
	err = gormConn.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	// Insert a new product
	newProduct := Product{
		Name:  "Milk",
		Price: 1.25,
	}
	tx := gormConn.Create(&newProduct)
	if tx.Error != nil {
		log.Fatalf("failed to create product: %v", tx.Error)
	}
	fmt.Printf("New product created with ID %d\n", newProduct.ID)

	// Find all products
	var products []Product
	tx = gormConn.Find(&products)
	if tx.Error != nil {
		log.Fatalf("failed to find products: %v", tx.Error)
	}
	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}

	// Close the database connection
	sqlDB, err := gormConn.DB()
	if err != nil {
		log.Fatalf("failed to get database object: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		fmt.Println("failed to close connection:", err)
	}
}

// Product model
type Product struct {
	ID    uint    `gorm:"column:id;primaryKey;autoIncrement"`
	Name  string  `gorm:"column:name;not null"`
	Price float64 `gorm:"column:price;not null"`
}

func (*Product) TableName() string {
	return "products"
}
