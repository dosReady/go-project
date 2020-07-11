package dao

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// postgres driver
	"github.com/dlog/core"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Setup export
func Setup() *gorm.DB {
	var err error
	dbMeta := core.GetConfig().DB
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbMeta.Host, dbMeta.Port, dbMeta.User, dbMeta.Password, dbMeta.Database)
	db, err = gorm.Open("postgres", dbinfo)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
