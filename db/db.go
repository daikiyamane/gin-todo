package db

import (
	"gin-todo/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

//InitDB inits DB
func InitDB() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	AutoMigration()
}

//GetDB returns *gorm.DB
func GetDB() *gorm.DB {
	return db
}

//CloseDB closed db
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

//AutoMigration does automigration
func AutoMigration() {
	db.AutoMigrate(&entity.Todo{})
}
