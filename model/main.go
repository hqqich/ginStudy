package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syndtr/goleveldb/leveldb"
	"jyksServer/common"
	"log"
)

// mysql 操作
var DB *gorm.DB

// levelDB 操作
var LevelDB *leveldb.DB

// InitDB 初始化DB
func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", common.DataSource)
	db.LogMode(true) // 打印sql
	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}

// InitDB 初始化DB
func InitLevelDB() (*leveldb.DB, error) {

	db, err := leveldb.OpenFile("./db", nil)

	if err == nil {
		LevelDB = db
		return LevelDB, err
	} else {
		log.Fatal(err)
	}

	return nil, err
}
