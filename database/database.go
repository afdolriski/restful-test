package database

import (
	"fmt"

	"github.com/afdolriski/restful-test/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // use sqlite dialect
)

// InitDB to initials the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/test.db")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to sqlite3 database")
	models.Migrate(db)
	return db, err
}
