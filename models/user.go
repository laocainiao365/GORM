package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Age int	`json:"age"`
	Sex string `json:"sex"`
}

func (User) TableName() string {
	return "user"
}

var DB *gorm.DB
func init() {
	DB,_ = gorm.Open("mysql", "root:12345678@/beego?charset=utf8&parseTime=True&loc=Local")
}
