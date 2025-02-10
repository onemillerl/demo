package dal

import (
	"gomall_demo/app/auth/biz/dal/mysql"
	"gomall_demo/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
