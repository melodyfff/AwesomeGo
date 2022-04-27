package orm

import "github.com/jinzhu/gorm"

// golang1.18 后可以这样使用
//func CreateTable(db *gorm.DB, value any) {
// 这里的any等同于： type any = interfance{}

func CreateTable(db *gorm.DB, value interface{}) {
	//判断表是否存在
	if db.HasTable(value) {
		// 自动迁移数据结构(table schema)
		// 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
		// db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
		db.AutoMigrate(value) //存在就自动适配表，也就说原先没字段的就增加字段
	} else {
		db.CreateTable(value) //不存在就创建新表
	}
}

func Insert(db *gorm.DB, value interface{}) {
	db.Create(value)
}
