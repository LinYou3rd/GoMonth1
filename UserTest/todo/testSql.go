package todo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "tianxiaDIYI427", "localhost", 3306, "testfortask")
	db, err = gorm.Open("mysql", constr)
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&todoModel{})
}
