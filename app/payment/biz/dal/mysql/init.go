package mysql

import (
	"fmt"
	"os"

	"gomall_demo/app/payment/biz/model"

	"gomall_demo/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.PaymentLog{}) // 自动注册

	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}
