package dal

import (
	"gomall_demo/app/email/biz/dal/mysql"
	"gomall_demo/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
