package consul

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"hat_push/tool/consul/dynamic"
	"hat_push/tool/consul/static"
	"hat_push/util"
)

var serviceInfo *consul.ServiceInfo

func InitConsul() {
	var err error
	if util.IsLocalRun() {
		serviceInfo = &consul.ServiceInfo{
			Cluster: util.GetLocalClusterArgs(),
			Service: util.GetLocalServerArgs(),
			Index:   util.GetLocalArgsIndex(),
		}
		if err = consul.WatchConfigByService(serviceInfo, static.NewStatic()); err != nil {
			klog.Panic(err)
		}
	} else {
		if err = consul.WatchConfig(static.NewStatic()); err != nil {
			klog.Panic(err)
		}
		if serviceInfo, err = consul.GetServiceInfoByPath(); err != nil {
			klog.Panic(err)
		}
	}
	initDynamic(serviceInfo)
}

// initDynamic 初始化动态配置
func initDynamic(ser *consul.ServiceInfo) {
	dynamicDir := ser.Cluster + "/" + ser.Service
	err := consul.WatchDir(dynamicDir, dynamic.GetDynamicDirCfg())
	if err != nil {
		klog.Panic(err)
	}
	dynamic.InitDynamicData(dynamicDir, ser)
}

func WhoAmI() *consul.ServiceInfo {
	return serviceInfo
}

func IsMasterIndex() bool {
	return serviceInfo.Index == dynamic.GetMasterNode()
}
