package static

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
)

var static = &StaticCfg{}

type StaticCfg struct {
	CommonDBUrl string // db地址
	CommonDb    string // db库名
	NotifyUrl   string //推送地址
}

func NewStatic() *StaticCfg {
	return static
}

func (this_ *StaticCfg) Reload() {
	static = this_
	klog.Info("reload consul 静态配置完成")
}

func (this_ *StaticCfg) New() consul.IConfig {
	return &StaticCfg{}
}

//==================== 调用函数 ====================//

//StaticRedisUrl redis地址
func StaticRedisUrl() string {
	return "static.Redis"
}

//StaticCommonDBUrl 公共数据db地址
func StaticCommonDBUrl() string {
	return static.CommonDBUrl
}

//StaticCommonDb
func StaticCommonDb() string {
	return static.CommonDb
}

func GetNotifyUrl() string {
	return static.NotifyUrl
}
