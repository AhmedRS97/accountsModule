package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"fmt"
)

var db *gorm.DB // the database

func init() {
	// this function will initialize the sqlite3 database
	// it will called automatically when loaded by go

	// below line will create and open a connection to the sqlite3 database file
	conn, err := gorm.Open("sqlite3", "database.db")

	// this will print any error that came from connecting to the database
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	// this will migrate the models
	db.Debug().AutoMigrate(&Account{})
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
