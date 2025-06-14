package initialziers

import (
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Connect2DB() {
	var err error
	dsn := os.Getenv("DBURL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect 2 DB")
	}
}
