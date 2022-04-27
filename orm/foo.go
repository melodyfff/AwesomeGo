package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// gorm.Model 的定义
//type Model struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}

type Foo struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string
	Sex   string `gorm:"type:enum('FeMale', 'Male');default:'Male'"`
}

// 设置表名，默认是结构体的名的复数形式

func (Foo) TableName() string {
	return "foo"
}

// 有点JAVA ToString的感觉

func (f Foo) String() string {
	return fmt.Sprintf("{id:%v , name:%v , sex:%v}", f.ID, f.Name, f.Sex)
}
