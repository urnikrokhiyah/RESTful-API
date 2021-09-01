package config

import (
	"mvc/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// config := map[string]string{
	// 	"DB_Username": "root",
	// 	"DB_Password": "alta123",
	// 	"DB_Host":     "localhost",
	// 	"DB_Port":     "3306",
	// 	"DB_Name":     "mvc",
	// }

	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config["DB_Username"],
	// 	config["DB_Password"],
	// 	config["DB_Host"],
	// 	config["DB_Port"],
	// 	config["DB_Name"],
	// )

	connection := os.Getenv("CONNECTION")

	var err error

	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
