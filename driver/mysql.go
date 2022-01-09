package driver

import (
	biData "invoice-api/features/billissuer/data"
	clData "invoice-api/features/client/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// For Linux
	dsn := "root:admin@tcp(127.0.0.1)/invoicein?parseTime=true"
	// For Windows
	// dsn := "root:@tcp(127.0.0.1)/moviein?parseTime=true"
	// For Amazon RDS
	// dsn := "admin:40fied40@tcp(moviein.c4v71mtnu5pg.us-east-2.rds.amazonaws.com)/moviein?parseTime=true"

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
	DB.AutoMigrate(&biData.BillIssuer{}, &clData.Client{})
}
