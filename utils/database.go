package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
func DB() (*gorm.DB) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		return nil
	} else {
		return db
	}
}