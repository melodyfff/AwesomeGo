package orm

import (
	"fmt"
	"testing"
)

func TestFooCrate(t *testing.T) {
	CreateTable(DB, &Foo{})
}

func TestFooInsert(t *testing.T) {
	Insert(DB, &Foo{Name: "ok", Email: "hello@qq.com", Sex: ""})
	Insert(DB, &Foo{Name: "ok2", Email: "hello@qq.com", Sex: "FeMale"})
	Insert(DB, &Foo{Name: "ok3", Email: "hello@qq.com", Sex: "FeMale"})
}

func TestFooQuery(t *testing.T) {
	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL ORDER BY `foo`.`id` ASC LIMIT 1
	q1 := DB.Debug().First(&Foo{})
	fmt.Println(q1.Value)

	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL ORDER BY `foo`.`id` DESC LIMIT 1
	q2 := DB.Debug().Last(&Foo{})
	fmt.Println(q2.Value)

	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL LIMIT 1
	q3 := DB.Debug().Take(&Foo{})
	fmt.Println(q3.Value)

	// 根据主键查询 SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL AND ((`foo`.`id` = 3)) ORDER BY `foo`.`id` ASC LIMIT 1
	q4 := DB.Debug().First(&Foo{}, 3)
	fmt.Println(q4.Value)

	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL AND ((name like '%ok%'))
	q5 := DB.Debug().Where("name like ?", "%ok%").Find(&[]Foo{})
	fmt.Println(q5.Value)

	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL AND ((name in ('ok','ok2')))
	q6 := DB.Debug().Where("name in (?)", []string{"ok", "ok2"}).Find(&[]Foo{})
	fmt.Println(q6.Value)

	// SELECT * FROM `foo`  WHERE `foo`.`deleted_at` IS NULL AND ((name='ok2')) ORDER BY `foo`.`id` ASC LIMIT 1
	q7 := DB.Debug().Where("name=?", "ok2").First(&Foo{})
	fmt.Println(q7.Value)

	var count int
	var tmp []Foo
	DB.Debug().Where("name=?", "ok2").First(&tmp).Count(&count)
	fmt.Println(count)
	fmt.Println(tmp)
}
