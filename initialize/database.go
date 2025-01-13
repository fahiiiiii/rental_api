// initialize/database.go
package initialize

import (
    "fmt"
    "log"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/beego/beego/v2/server/web"
)

var DB *gorm.DB

func InitializeDB() {
    // Fetch database configurations from app.conf
    driver, _ := web.AppConfig.String("db::driver")
    user, _ := web.AppConfig.String("db::user")
    password, _ := web.AppConfig.String("db::password")
    host, _ := web.AppConfig.String("db::host")
    port, _ := web.AppConfig.String("db::port")
    name, _ := web.AppConfig.String("db::name")

    // Construct the DSN (Data Source Name)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, name, port)

    // Initialize the database connection
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto-migrate your models here if needed
    // DB.AutoMigrate(&models.User{}, &models.Product{})

    log.Println("Database connected successfully!")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
    return DB
}
