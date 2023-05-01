package initializers

import (
	"log"
	"os"

	"github.com/AdityaRajSingh/go-blog-post/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is a global variable that represents the database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(config *Config) {
	// Connect to the database
	var err error

	dsn := config.DSN
	log.Println("Connecting to the Database... ")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", dsn+err.Error())
		os.Exit(1)
	}
	log.Println("🚀 Connected Successfully to the Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&models.BlogPost{})
	DB = db
}
