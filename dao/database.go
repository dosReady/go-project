package dao

import (
	"fmt"
	"log"
	"time"

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

// CommonModel : 공통 테이블 모델
type CommonModel struct {
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

// TbUser : 유저 테이블 모델
type TbUser struct {
	LoginID     string `gorm:"type:varchar(100);"`
	Password    string `gorm:"type:varchar(100);not null;"`
	Role        string `gorm:"varchar(100);not null;"`
	AccessToken string `gorm:"text;"`
	CommonModel
}

// TbPost : Post 테이블 모델
type TbPost struct {
	PostKey      string `gorm:"type:bigserial;primary_key;auto_increment;"`
	PostTitle    string `gorm:"type:varchar(100);not null;"`
	PostSubTitle string `gorm:"varchar(100);not null;"`
	PostContent  string `gorm:"text;"`
	PostCategory string `gorm:"varchar(100);not null;"`
	CommonModel
}
