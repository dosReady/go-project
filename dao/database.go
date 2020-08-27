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

type Session struct {
	Db     *gorm.DB
	isTran bool
}

// Setup export
func Setup(pIsTran bool) (session *Session) {
	var db *gorm.DB

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
	session = &Session{
		isTran: pIsTran,
	}
	if pIsTran {
		db = db.Begin()
	}

	session.Db = db
	return session
}

// Close export
func (session *Session) Close() {
	if session.isTran {
		if r := recover(); r != nil {
			session.Db.Rollback()
		} else {
			session.Db.Commit()
		}
	}

	session.Db.Close()
}

// InitializeDB: server start 시 미들웨어 작동
func (session *Session) InitializeDB() {
	log.Println("========================= DB 초기화 ==========================")
	defer session.Close()

	db := session.Db
	db.AutoMigrate(TbUser{}, TbPost{}, TbTag{}, TbTagMap{})
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

// TbTag : Tag 테이블 모델
type TbTag struct {
	TagKey  string `gorm:"type:bigserial;primary_key;auto_increment;"`
	TagName string `gorm:"type:varchar(100);not null"`
	CommonModel
}

// TbTagMap: Post에 연결된 Tag 맵 테이블
type TbTagMap struct {
	PostKey string `gorm:"type:bigserial;not null;"`
	TagKey  string `gorm:"type:bigserial;not null;"`
	CommonModel
}
