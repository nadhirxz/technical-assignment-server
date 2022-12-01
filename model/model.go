package model

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error

	db, err = gorm.Open(sqlite.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})

	if err != nil {
		fmt.Println(err)
	}
}
