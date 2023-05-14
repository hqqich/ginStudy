package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
)

// sqlite entity
//type User struct {
//	Id          int    `json:"id"`
//	Username    string `json:"username" gorm:"unique;type:string"`
//	Password    string `json:"password" gorm:"not null;type:string;"`
//	DisplayName string `json:"displayName" gorm:"type:string;"`
//	Role        int    `json:"role" gorm:"type:int;default:1"`   // admin, common
//	Status      int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
//}

// mysql entity
type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username" gorm:"unique;type:varchar(255)"`
	Password    string `json:"password" gorm:"not null;type:varchar(255);"`
	DisplayName string `json:"displayName" gorm:"type:varchar(255);"`
	Role        int    `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status      int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
}

func (user *User) Insert() error {
	var err error
	err = DB.Create(user).Error
	return err
}

func (user *User) Update() error {
	var err error
	err = DB.Model(user).Updates(user).Error
	return err
}

func (user *User) Delete() error {
	var err error
	err = DB.Delete(user).Error
	return err
}

func (user *User) ValidateAndFill() {
	// When querying with struct, GORM will only query with non-zero fields,
	// 当使用struct查询时，GORM将只查询非零字段，
	// that means if your field’s value is 0, '', false or other zero values,
	// 这意味着如果你的字段值为0，"，false或其他零值，
	// it won’t be used to build query conditions
	// 它不会被用来构建查询条件
	DB.Where(&user).First(&user)
}

// Validate 返回是否能登录，true代表能登录，false代表不能登录
func (user *User) Validate() bool {
	username := user.Username
	//password := user.Password

	contains := strings.Contains(username, "admin") || strings.Contains(username, "tsinglink")

	return contains
}
