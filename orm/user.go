package orm

import (
	"database/sql"
	"github.com/jinzhu/gorm"
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

type (
	User struct {
		gorm.Model
		UserId      int64 `gorm:"index"` //设置一个普通的索引，没有设置索引名，gorm会自动命名
		Birthday    time.Time
		Age         int           `gorm:"column:age"`                     //column:一个tag，可以设置列名称
		Name        string        `gorm:"size:255;index:idx_name_add_id"` //size:设置长度大小，index:设置索引，这个就取了一个索引名
		Num         int           `gorm:"AUTO_INCREMENT"`
		Email       string        `gorm:"type:varchar(100);unique_index"` //type:定义字段类型和大小
		AddressID   sql.NullInt64 `gorm:"index:idx_name_add_id"`
		IgnoreMe    int           `gorm:"_"`
		Description string        `gorm:"size:2019;comment:'用户描述字段'"` //comment：字段注释
		Status      string        `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"`
	}
)

// 设置表名，默认是结构体的名的复数形式

func (User) TableName() string {
	return "VIP_USER"
}
