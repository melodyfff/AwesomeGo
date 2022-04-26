package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // "_" 引入后面的包名 而不直接使用里面的定义的函数、变量、资源等
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// gorm模型定义 https://gorm.io/zh_CN/docs/models.html

// gorm.Model 的定义
//type Model struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}

type User struct {
	gorm.Model
	UserId      int64 `gorm:"index"`
	Birtheday   time.Time
	Age         int           `gorm:"column:age"`
	Name        string        `gorm:"size:255;index:idx_name_add_id"`
	Num         int           `gorm:"AUTO_INCREMENT"`
	Email       string        `gorm:"type:varchar(100);unique_index"`
	AddressID   sql.NullInt64 `gorm:"index:idx_name_add_id"`
	IgnoreMe    int           `gorm:"_"`
	Description string        `gorm:"size:2019;comment:'用户描述字段'"`
	Status      string        `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"`
}

// 设置表名，默认是结构体的名的复数形式

func (User) TableName() string {
	return "VIP_USER"
}

func create(db *gorm.DB, value interface{}) {
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

func main() {
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("connect db err: ", err)
	}
	defer db.Close()

	create(db, &User{})
}
