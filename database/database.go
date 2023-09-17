package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type database struct {
	HOST     string
	USER     string
	PASSWORD string
	NAME     string
}

func NewDb() database {
	db := database{
		HOST:     "localhost",
		USER:     "postgres",
		PASSWORD: "1234",
		NAME:     "postgres",
	}
	return db
}

func GetConnection() *gorm.DB {
	db := NewDb()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", db.HOST, db.USER, db.PASSWORD, db.NAME)

	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	Db.Logger = logger.Default.LogMode(logger.Info)

	return Db
}

func CreateItem(model interface{}) {
	db := GetConnection()
	db.Create(model)

}

func GetItems(model interface{}) {}

func GetItem(model interface{}) {}

func UpdateItem(model interface{}) {}

func DeleteItem(model interface{}) {}
