package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	Init()
}

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gin-gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDatabase() *gorm.DB {
	return DB
}
