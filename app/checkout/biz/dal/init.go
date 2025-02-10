package dal

import (
	"gomall_demo/app/checkout/biz/dal/mysql"
	"gomall_demo/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
