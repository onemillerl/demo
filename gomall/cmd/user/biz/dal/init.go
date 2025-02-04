package dal

import (
	"github.com/onemillerl/gomall/biz/dal/mysql"
	"github.com/onemillerl/gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
