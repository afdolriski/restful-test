package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate the models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Image{})

	fmt.Println("Migration Success")
}
