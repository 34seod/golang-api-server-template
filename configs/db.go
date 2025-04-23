package configs

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	time.Sleep(time.Second * 2)
	cnf := Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.DBUserName, cnf.DBPassword, cnf.DBHost, cnf.DBPort, cnf.DBName)

	connect(dsn)
}

// only for test
func ConnectTestDB() {
	cnf := Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"test", "test", "localhost", cnf.DBPort, "test")

	connect(dsn)
}

func connect(dsn string) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("fail to connect DB")
	}

	DB = database
}
