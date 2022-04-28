package orm

import (
	"testing"

	//_ "github.com/go-sql-driver/mysql" // "_" 引入后面的包名 而不直接使用里面的定义的函数、变量、资源等
	// 这里两个dialects导入方式本质是一样的，可以观察github.com/jinzhu/gorm/dialects/ 目录下获取更多支持
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestSetTable(t *testing.T) {
	//db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	log.Println("connect db err: ", err)
	//}
	//defer db.Close()

	CreateTable(DB, &User{})
}
