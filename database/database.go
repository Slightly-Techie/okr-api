package database

import (
	"fmt"
	"log"

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

func NewDb() *database {
	db := &database{
		HOST:     "localhost",
		USER:     "postgres",
		PASSWORD: "1234",
		NAME:     "postgres",
	}
	return db
}

var db *gorm.DB

func InitDB() error {
	Db := NewDb()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", Db.HOST, Db.USER, Db.PASSWORD, Db.NAME)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	log.Println("Connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil. Make sure to call InitDB() first.")
	}
	return db
}

func CreateItem(model interface{}) error {
	if db == nil {
		return fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}

	result := db.Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetItems[T any](model *[]T) (*[]T, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}

	if err := db.Find(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func GetItem(model interface{}) {}

func UpdateItem(model interface{}) {}

func DeleteItem(model interface{}) {}
