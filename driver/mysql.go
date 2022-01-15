package driver

import (
	"fmt"
	"invoice-api/config"
	biData "invoice-api/features/billissuer/data"
	bidData "invoice-api/features/billissuerdetail/data"
	clData "invoice-api/features/client/data"
	inData "invoice-api/features/invoice/data"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println("Config : ", config)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	// For Linux
	// dsn := "root:admin@tcp(127.0.0.1)/invoicein?parseTime=true"
	// For Windows
	// dsn := "root:@tcp(127.0.0.1)/invoice?parseTime=true"
	// For Amazon RDS
	// dsn := "admin:40fied40@tcp(moviein.c4v71mtnu5pg.us-east-2.rds.amazonaws.com)/moviein?parseTime=true"

	// var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db

	DB.AutoMigrate(&biData.BillIssuer{}, &clData.Client{}, &inData.Invoice{}, &bidData.BillIssuerDetail{})
}
