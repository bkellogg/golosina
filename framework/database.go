package framework

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// We include this package to be able to connect to mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database is our wrapper for the gorm.DB
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database and gives us
// the connected database pointer for queries
func NewDatabase() *Database {

	// Sometimes we don't need/have a database so
	// we will only use it whenever we need to
	if os.Getenv("DB_ENABLED") != "true" {
		return nil
	}

	dbConnString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	log.Println("Trying to connect to DB")
	log.Println(dbConnString)

	db, err := gorm.Open(os.Getenv("DB_DRIVER"), dbConnString)
	if err != nil {
		log.Println(err)
		panic("Failed to connect to Database")
	}

	return &Database{db}
}
