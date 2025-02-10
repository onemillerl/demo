package dal

import (
	"gomall_demo/app/payment/biz/dal/mysql"
	"gomall_demo/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
