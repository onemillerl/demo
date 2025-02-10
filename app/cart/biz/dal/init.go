package dal

import (
	"gomall_demo/app/cart/biz/dal/mysql"
	"gomall_demo/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
