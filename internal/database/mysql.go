package database

import (
	"evermosTest/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strconv"
	"time"
)

func NewMySQL() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxRetry := 10

	for i := 0; i < maxRetry; i++ {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //mengeprint query yg digunakan
		})

		if i > 0 {
			fmt.Println("DB Connection : Retry Mechanism [" + strconv.Itoa(i) + "x]")
		}
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	s, err := db.DB()
	if err = s.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return db, nil
}
