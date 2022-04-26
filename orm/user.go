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
		UserId      int64 `gorm:"index"`
		Birthday    time.Time
		Age         int           `gorm:"column:age"`
		Name        string        `gorm:"size:255;index:idx_name_add_id"`
		Num         int           `gorm:"AUTO_INCREMENT"`
		Email       string        `gorm:"type:varchar(100);unique_index"`
		AddressID   sql.NullInt64 `gorm:"index:idx_name_add_id"`
		IgnoreMe    int           `gorm:"_"`
		Description string        `gorm:"size:2019;comment:'用户描述字段'"`
		Status      string        `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"`
	}
)

// 设置表名，默认是结构体的名的复数形式

func (User) TableName() string {
	return "VIP_USER"
}
