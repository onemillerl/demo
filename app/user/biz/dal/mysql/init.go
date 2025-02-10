package mysql

import (
	"fmt"
	"os"

	"gomall_demo/app/user/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// func Init() {
// 	// conf.GetConf().MySQL.DSN
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("MYSQL_USER"),
// 		os.Getenv("MYSQL_PASSWORD"), // 修正拼写错误
// 		os.Getenv("MYSQL_HOST"),
// 		os.Getenv("MYSQL_DATABASE"), // 修正拼写错误
// 	)
// 	// fmt.Println("DSN:", dsn) // 打印 dsn 确认是否正确

// 	DB, err = gorm.Open(mysql.Open(dsn),
// 		&gorm.Config{
// 			PrepareStmt:            true,
// 			SkipDefaultTransaction: true,
// 		},
// 	)
// 	DB.AutoMigrate(&model.User{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := DB.Use(tracing.NewPlugin()); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("MySQL connection established", dsn)
// }

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"), // 修正拼写错误
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"), // 修正拼写错误
	)
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.User{})
		DB.AutoMigrate( //nolint:errcheck
			&model.User{},
			&model.ProductId{}, // 确保 ProductId 也被迁移
		)
		if needDemoData {
			DB.Exec("INSERT INTO `user` (`id`,`created_at`,`updated_at`,`email`,`password_hashed`) VALUES (1,'2023-12-26 09:46:19.852','2023-12-26 09:46:19.852','123@admin.com','$2a$10$jTvUFh7Z8Kw0hLV8WrAws.PRQTeuH4gopJ7ZMoiFvwhhz5Vw.bj7C')")
		}
	}
}
