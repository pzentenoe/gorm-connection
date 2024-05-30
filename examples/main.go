package main

import (
	"fmt"
	"github.com/pzentenoe/gorm-connection/sql"
	"log"
	"time"
)

func main() {
	connection, err := sql.NewSQLConnection(sql.Config().
		SetSQLDialect(sql.Postgres).
		Host("localhost").
		DatabaseName("test_db").
		User("test_user").
		Password("test_password").
		Port(sql.Postgres.DefaultPort).    //Optional because there is a default port in sql.Postgres.DefaultPort
		Timezone("America/Santiago").      //Optional default value "UTC"
		MaxIdleConns(5).                   //Optional value default value 10
		MaxOpenConns(50).                  // Optional option default value Optional option default value
		ConnMaxLifetime(time.Minute * 60). // Optional option default value time.Minute * 30
		ConnMaxIdleTime(50),               // Optional option default value 100
	)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//Migrate
	gormConn, _ := connection.GetConnection()
	err = gormConn.AutoMigrate(Product{})
	if err != nil {
		return
	}

	//Find by id
	var product Product
	tx := gormConn.Find(&product, "id = ?", 1)
	if tx.Error != nil {
		log.Fatalf("failed to find product: %v", tx.Error)
	} else {
		fmt.Println("product found:", product)
	}

	if product.ID != 0 {
		//Insert
		newProduct := Product{
			Name:  "Milk",
			Price: 1.25,
		}
		tx = gormConn.Create(&newProduct)
		if tx.Error != nil {
			log.Fatalf("failed to create product: %v", tx.Error)
		}
		fmt.Printf("new product created id %d\n", newProduct.ID)
	}

	productRepository := NewProductRepository(connection)

	//Find
	products, err := productRepository.FindAllProducts()
	if err != nil {
		return
	}
	//Do it anything with products array
	for _, product := range products {
		fmt.Print(product)
	}
	db, _ := gormConn.DB()
	err = db.Close()
	if err != nil {
		fmt.Println("failed to close connection ", err)
	}
}

// Example
// Use your created connection for create your repository
type ProductRepository struct {
	dbConnection sql.Connection
}

func NewProductRepository(dbConnection sql.Connection) *ProductRepository {
	return &ProductRepository{dbConnection: dbConnection}
}

func (r *ProductRepository) FindAllProducts() ([]*Product, error) {
	db, err := r.dbConnection.GetConnection()
	if err != nil {
		return nil, err
	}
	products := make([]*Product, 0)
	dbResponse := db.Find(&products)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return products, nil
}

type Product struct {
	ID    uint    `gorm:"column:id; primaryKey; autoIncrement;"`
	Name  string  `gorm:"column:name; not null"`
	Price float64 `gorm:"column:price; not null"`
}

func (*Product) TableName() string {
	return "products"
}
