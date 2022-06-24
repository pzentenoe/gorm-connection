# gorm-connection

Gorm Connection provides a wrapper for use gorm easily

### Use mode
```go
import (
    sql "github.com/pzentenoe/gorm-connection/sql"
)


func main(){
    connection := sql.NewSQLConnection(sql.Config().
    SetSQLDialect(sql.MySQL).
    Host("localhost").
    Port("3306").
    DatabaseName("database_name").
    User("user").
    Password("password"),
    )

    productRepository := NewProductRepository(connection)
    // method call example, you should create your our methods
	productRepository.GetProducts()
}

//Example
// Use your created connection for create your repository
type productRepository struct{
    dbConnection sql.Connection
}

func NewProductRepository(connection sql.Connection)*productRepository{
    return &productRepository{dbConnection: connection}
}

func (r *productRepository) FindAllProducts() ([]*Product, error) {
    db := r.dbConnection.GetConnection()
	products := make([]*Product, 0)
    dbResponse := db.Find(&products)
    if dbResponse.Error != nil {
        return nil, dbResponse.Error
    }
    return products, nil
}

type Product struct {
    ID        string `gorm:"column:id; primaryKey;"`
    Name     string `gorm:"column:name; not null"`
    Price  string `gorm:"column:price; not null"`
}

func (*User) TableName() string {
    return "products"
}


```

### Author
[Pablo Zenteno](https://github.com/pzentenoe)
