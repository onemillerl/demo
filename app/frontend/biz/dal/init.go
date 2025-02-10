package dal

import (
	"gomall_demo/app/frontend/biz/dal/redis"

	"gomall_demo/app/frontend/biz/dal/mysql"
)

func Init() {
	redis.Init()
	mysql.Init()
}
