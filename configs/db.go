package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	connect := "root:root@tcp(127.0.0.1:8889)/db_luxspace?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
