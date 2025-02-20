package main

import (
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/configs"
	"hat_push/api"
	"hat_push/core"
	"hat_push/pbgo/pbpush"
	"hat_push/tool"
	"hat_push/util"
)

func init() {
	util.InitUtil()
	tool.InitTool()
	core.InitCore()
	klog.Info("server init ok")
}

func main() {
	klog.Info("start kite")
	if configs.Upgrade() {
		// 热更新这里不处理
	} else {
	}
	startKite()
}

func startKite() {
	if util.IsLocalRun() {
		configs.GlobalCommonObject.MetricsReportURL = ""
		kite.WhoAmI(util.GetLocalClusterArgs(), util.GetLocalServerArgs(), util.GetLocalArgsIndex())
	}
	pbpush.RegPushApiServer(&api.PushApi{})
	kite.Serve(&Process{})
}

type Process struct {
}

func (this_ *Process) OldBefore(upGradeLevel int) {
}

func (this_ *Process) NewBefore(upGradeLevel int) {
}

func (this_ *Process) OldAfter(upGradeLevel int) {
}

func (this_ *Process) NewAfter(upGradeLevel int) {
}

func (this_ *Process) SendData(send func(data []byte)) error {
	return nil
}

func (this_ *Process) ReceiveData(data []byte) error {
	return nil
}

func (this_ *Process) Stop() {
	AfterStopServe()
}

func AfterStopServe() {
	tool.ReleaseTool()
}
