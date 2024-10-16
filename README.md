# gorm-connection

## Overview

Gorm Connection provides a convenient wrapper to simplify interactions with GORM, reducing repetitive tasks and streamlining the process of database connectivity and operations.

[![codecov](https://codecov.io/github/pzentenoe/gorm-connection/graph/badge.svg?token=3W164MZ18S)](https://codecov.io/github/pzentenoe/gorm-connection)
![CI](https://github.com/pzentenoe/gorm-connection/actions/workflows/actions.yml/badge.svg)
![GoDoc](https://github.com/pzentenoe/gorm-connection/actions/workflows/documentation.yml/badge.svg)
![Quality Gate](https://sonarqube.vikingcode.cl/api/project_badges/measure?project=gorm-connection&metric=alert_status&token=sqb_92a3538dbd1a098295e4de2086d6ab11c0243ad9)
![Coverage](https://sonarqube.vikingcode.cl/api/project_badges/measure?project=gorm-connection&metric=coverage&token=sqb_92a3538dbd1a098295e4de2086d6ab11c0243ad9)
![Bugs](https://sonarqube.vikingcode.cl/api/project_badges/measure?project=gorm-connection&metric=bugs&token=sqb_92a3538dbd1a098295e4de2086d6ab11c0243ad9)

## Features

## Features

- **Simplified Configuration:** Facilitates connection management with GORM, reducing initial setup complexity.
- **Multiple SQL Dialect Support:** Supports common SQL dialects, including SQL Server, MySQL, PostgreSQL, and SQLite.
- **Sensible Default Values:** Provides default configurations for SQL dialects and their ports, making connections easier.
- **Easy Integration:** Simple to integrate into existing projects with minimal additional configuration.
- **Built-in Pagination:** Includes a pagination function to simplify database queries with paginated results.

## Installation

You can install the package using `go get`:

```bash
go get github.com/pzentenoe/gorm-connection/sql
```


## Go Doc
<a href="https://pzentenoe.github.io/gorm-connection" target="_blank">godoc</a>


## Suported Dialects

* SQLServer
* MySQL
* Postgres
* SQLite


### Support Me
If this project has been helpful to you, consider supporting me:

<a href="https://www.buymeacoffee.com/pzentenoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

Thank you for your support! ❤️

## Usage
Below are examples of how to use GORM Connection with different databases.

### Connecting to SQLite
```go
package main

import (
    "fmt"
    "log"

    "github.com/pzentenoe/gorm-connection/sql"
    "gorm.io/gorm"
)

func main() {
    // Create a connection to SQLite
    connection, err := sql.NewSQLiteConnection("test_db")
    if err != nil {
        log.Fatalf("Failed to connect to SQLite database: %v", err)
    }

    // Get GORM connection
    gormConn, err := connection.GetConnection()
    if err != nil {
        log.Fatalf("Failed to get GORM connection: %v", err)
    }

    // Insert a new product
    newProduct := Product{
        Name:  "Milk",
        Price: 1.25,
    }
    tx := gormConn.Create(&newProduct)
    if tx.Error != nil {
        log.Fatalf("Failed to create product: %v", tx.Error)
    }
    fmt.Printf("New product created with ID %d\n", newProduct.ID)

    // Query products
    var products []Product
    tx = gormConn.Find(&products)
    if tx.Error != nil {
        log.Fatalf("Failed to query products: %v", tx.Error)
    }
    for _, product := range products {
        fmt.Printf("Product: %+v\n", product)
    }
}

// Product model
type Product struct {
    ID    uint    `gorm:"primaryKey"`
    Name  string  `gorm:"not null"`
    Price float64 `gorm:"not null"`
}

```
### Connecting to PostgreSQL
MySQL an SQLServer are similar to PostgresSQL
```go
package main

import (
    "fmt"
    "log"

    "github.com/pzentenoe/gorm-connection/sql"
    "gorm.io/gorm/logger"
)

func main() {
    // Create a connection to PostgreSQL
    connection, err := sql.NewPostgresConnection(
        "localhost",
        "test_db",
        "user",
        "password",
        sql.WithPort(5432),
        sql.WithLogLevel(logger.Info),
    )
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
    }

    // Get GORM connection
    gormConn, err := connection.GetConnection()
    if err != nil {
        log.Fatalf("Failed to get GORM connection: %v", err)
    }

    // Insert a new product
    newProduct := Product{
        Name:  "Bread",
        Price: 0.80,
    }
    tx := gormConn.Create(&newProduct)
    if tx.Error != nil {
        log.Fatalf("Failed to create product: %v", tx.Error)
    }
    fmt.Printf("New product created with ID %d\n", newProduct.ID)

    // Query products with pagination
    var products []Product
    page := 1
    pageSize := 10
    tx = gormConn.Scopes(sql.Paginate(page, pageSize)).Find(&products)
    if tx.Error != nil {
        log.Fatalf("Failed to query products: %v", tx.Error)
    }
    for _, product := range products {
        fmt.Printf("Product: %+v\n", product)
    }
}

// Product model
type Product struct {
    ID    uint    `gorm:"primaryKey"`
    Name  string  `gorm:"not null"`
    Price float64 `gorm:"not null"`
}


```


### Examples

You can find example usages of this library in the [examples](https://github.com/pzentenoe/gorm-connection/tree/main/examples) folder of this repository.


## Contributing
Contributions are welcome! Please follow these steps:

1. Fork the project.
2. Create a branch for your feature (git checkout -b feature/new-feature).
3. Make your changes and add tests if necessary.
4. Ensure all tests pass (go test ./...).
5. Submit a pull request to the main branch.

## Testing

Execute the tests with:

```bash
go test ./...
```

## License
This project is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Changelog
For a detailed changelog, refer to [CHANGELOG.md](CHANGELOG.md).


### Author
[Pablo Zenteno](https://github.com/pzentenoe)
