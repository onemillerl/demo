package dal

import (
	"gomall_demo/app/product/biz/dal/mysql"
	"gomall_demo/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
