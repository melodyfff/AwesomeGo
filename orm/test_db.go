package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
)

var (
	DB *gorm.DB
)

const (
	MysqlDSN = "root:123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var err error
	// DSN(数据源名称)
	// 这里可以通过命令行设置参数来进行切换测试，目前先暂时写死MYSQL
	os.Setenv("DSN", MysqlDSN)
	os.Setenv("DIALECT", "mysql")

	if DB, err = OpenTestConnection(); err != nil {
		// 终止程序运行
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("DSN")
	switch os.Getenv("DIALECT") {
	case "mysql":
		fmt.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = MysqlDSN
		}
		db, err = gorm.Open("mysql", dbDSN)
	case "postgres":
		fmt.Println("testing postgres...")
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
		}
		db, err = gorm.Open("postgres", dbDSN)
	case "mssql":
		// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		// CREATE DATABASE gorm;
		// USE gorm;
		// CREATE USER gorm FROM LOGIN gorm;
		// sp_changedbowner 'gorm';
		fmt.Println("testing mssql...")
		if dbDSN == "" {
			dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		}
		db, err = gorm.Open("mssql", dbDSN)
	default:
		fmt.Println("testing sqlite3...")
		db, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
	}

	// db.SetLogger(Logger{log.New(os.Stdout, "\r\n", 0)})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.LogMode(true)
	} else if debug == "false" {
		db.LogMode(false)
	}

	// 连接池设置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return
}
