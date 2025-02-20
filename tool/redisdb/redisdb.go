package redisdb

import (
	"git.dhgames.cn/svr_comm/gcore/gredis"
	"hat_push/tool/consul/static"
)

func InitRedisDb() {
	gredis.Init(static.StaticRedisUrl())
}
