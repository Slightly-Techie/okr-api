package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
		HOST:     os.Getenv("POSTGRES_HOST"),
		USER:     os.Getenv("POSTGRES_USER"),
		PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		NAME:     os.Getenv("POSTGRES_DB"),
	}
	return db
}

var db *gorm.DB

func InitDB() error {
	er := godotenv.Load()
	if er != nil {
		log.Fatalf("Error loading .env file: %v", er)
	}
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

func CreateItem[T any](model T) error {
	if db == nil {
		return fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}

	result := db.Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetItems[T any](model *[]T, index, id string) ([]T, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}

	if err := db.Where(index+" = ?", id).Find(&model).Error; err != nil {
		return nil, err
	}

	return *model, nil
}

func GetItem[T any](model T, id string) error {
	if err := db.First(model, id).Error; err != nil {
		return err
	}
	return nil
}

func GetItemByEmail[T any](model T, email string) error {
	if err := db.First(model, "email = ?", email).Error; err != nil {
		return err
	}
	return nil
}

func UpdateItem[T any](model T, updates map[string]interface{}, index, id string) error {
	if db == nil {
		return fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}

	if err := db.First(&model, index+"= ?", id).Error; err != nil {
		return err
	}

	if err := db.Model(&model).Where(index+"= ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem[T any](model T, column, value string) error {
	if db == nil {
		return fmt.Errorf("database connection is nil. Make sure to call InitDB() first")
	}
	if err := db.Where(column+"= ?", value).Delete(&model).Error; err != nil {
		return err
	}

	return nil
}
