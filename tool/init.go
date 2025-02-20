package tool

import (
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"hat_push/tool/ants"
	"hat_push/tool/comDb"
	"hat_push/tool/consul"
	"hat_push/tool/geolite2"
	"hat_push/tool/goRuntime"
	"hat_push/tool/restyHttp"
)

func InitTool() {
	goRuntime.InitGoRuntime()
	consul.InitConsul()
	//redisdb.InitRedisDb()
	restyHttp.InitResty()
	comDb.InitComDb()
	ants.InitAnts()
	geolite2.InitGeoIp()
}

func ReleaseTool() {
	goRuntime.CloseGoRuntime()
	ants.ReleaseAnts()
	geolite2.CloseGeoIp()
	klog.Info("Tool Close OK")
}
