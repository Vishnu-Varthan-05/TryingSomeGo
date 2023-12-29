package database

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB()(*gorm.DB, error){
	dsn := "root:31125@tcp(127.0.0.1:3306)/best_shop?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, err
	}
	DB = db
	return db, nil
}
