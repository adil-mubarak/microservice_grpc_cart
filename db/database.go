package db

import (
	"log"
	"microservice_grpc_cart/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabse() (*gorm.DB, error) {
	dsn := "root:kl18jda183079@tcp(127.0.0.1:3306)/carts_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect database: %v",err)
		return nil,err
	}

	err = db.AutoMigrate(&models.Cart{})
	if err != nil{
		log.Fatalf("Failed to migrate table: %v",err)
		return nil,err
	}

	return db,nil
	
}
