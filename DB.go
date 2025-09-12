package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	Shortend  string `gorm:"uniqueIndex"`
	Long      string
	CreatorID int
}

type Users struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

type Server struct {
	db *gorm.DB
}

func DBinit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Urls{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
