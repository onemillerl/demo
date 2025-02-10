package dal

import (
	"gomall_demo/app/order/biz/dal/mysql"
	"gomall_demo/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
