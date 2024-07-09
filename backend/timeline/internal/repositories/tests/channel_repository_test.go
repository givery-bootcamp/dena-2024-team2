package repositories_test

import (
	"fmt"
	"myapp/internal/config"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	username := config.DBUserName
	if config.DBPassword != "" {
		username += ":" + config.DBPassword
	}
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tx := db.Begin()
	return tx
}

func TestChannelRepository(t *testing.T) {
	tx := SetupTestDB()
	tx.Omit("DeletedAt").Create(&post)
	tx.Rollback()
}
