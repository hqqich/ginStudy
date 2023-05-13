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
		db.AutoMigrate(&User{})
		createAdminAccount()
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

// 初始化用户
func createAdminAccount() {
	var user User
	DB.Where(User{Role: common.RoleAdminUser}).Attrs(User{
		Username:    "hqqich",
		Password:    "123456",
		Role:        common.RoleAdminUser,
		Status:      common.UserStatusEnabled,
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}
