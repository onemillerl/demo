package dal

import (
	"gomall_demo/app/user/biz/dal/mysql"
	"gomall_demo/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}


