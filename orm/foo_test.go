package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

var (
	TestDB *gorm.DB
)

func init() {
	// 初始化db
	// 创建gorm.DB对象的时候连接并没有被创建，在具体使用的时候才会创建。gorm内部，准确的说是database/sql内部会维护一个连接池，可以通过参数设置最大空闲连接数，连接最大空闲时间等。使用者不需要管连接的创建和关闭。
	if db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	} else {
		TestDB = db
	}
}

func TestFooCrate(t *testing.T) {
	CreateTable(TestDB, &Foo{})
}

func TestFooInsert(t *testing.T) {
	Insert(TestDB, &Foo{Name: "ok", Email: "hello@qq.com", Sex: ""})
	Insert(TestDB, &Foo{Name: "ok2", Email: "hello@qq.com", Sex: "FeMale"})
}

func TestFooQuery(t *testing.T) {
	q1 := TestDB.First(&Foo{})
	fmt.Println(q1.Value)

	q2 := TestDB.Where("name=?", "ok2").First(&Foo{})
	fmt.Println(q2.Value)
}
