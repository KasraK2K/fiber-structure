package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Connection struct {
	Conn *gorm.DB
}

func Connect() Connection {
	dsn := "host=localhost user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return Connection{Conn: db}
}

func Migrate(modelStruct interface{}) {
	var db = Connect().Conn

	//Create Table If Not Exist
	err := db.AutoMigrate(&modelStruct)
	if err != nil {
		return
	}
}
