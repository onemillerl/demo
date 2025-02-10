package mysql

import (
	"fmt"
	"os"

	"gomall_demo/app/auth/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"), // 修正拼写错误
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"), // 修正拼写错误
	)
	// fmt.Println("DSN:", dsn) // 打印 dsn 确认是否正确

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Authuser{})
	if err != nil {
		panic(err)
	}
}
