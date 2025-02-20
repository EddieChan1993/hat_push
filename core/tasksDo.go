package core

import (
	"context"
	"git.dhgames.cn/svr_comm/gcore/glog"
	"hat_push/cst"
	"hat_push/pbgo/pbpush"
	"hat_push/tool/consul/static"
	"hat_push/tool/goRuntime"
	"hat_push/tool/restyHttp"
	"time"
)

/*
*
推送任务分配处理
*/
var taskWorkers []*taskCh
var workersNums int32 = 3 //进程数
const doTaskdefTick = 5   //单位s
const buffer = 10
const msgNumsOk = 500 //单次处理任务数量

type taskCh struct {
	taskCh  chan *pbpush.PushUserTask
	allData []*pbpush.PushUserTask
}

func initTaskWorkers() {
	for i := int32(0); i < workersNums; i++ {
		tmp := make(chan *pbpush.PushUserTask, buffer)
		taskWorkers = append(taskWorkers, &taskCh{
			taskCh:  tmp,
			allData: make([]*pbpush.PushUserTask, 0, msgNumsOk),
		})
	}
	for _, worker := range taskWorkers {
		worker.run()
	}
}

func (t *taskCh) run() {
	goRuntime.GoRun(func(ctx context.Context) {
		ticker := time.NewTicker(doTaskdefTick * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				t.doAllTasks()
				return
			case <-ticker.C:
				t.doAllTasks()
			case data := <-t.taskCh:
				t.addTask(data)
			}
		}
	})
}

// doAllTasks 处理所有事件
func (t *taskCh) doAllTasks() {
	if len(t.allData) == 0 {
		return
	}
	tmpAllData := t.allData
	t.allData = t.allData[:0]
	glog.Infof("send start tasks %d", len(tmpAllData))
	var langMsg string
	bodyData := make(map[string]*cst.NotifyT)
	for _, per := range tmpAllData {
		langMsg = cst.GetLangMsg(per.HandleType, per.Language)
		if langMsg == "" {
			//glog.Warnf("GetLangMsg no type %d lang %s", per.HandleType, per.Language)
			continue
		}
		if notifyList, has := bodyData[per.Bundle]; has {
			notifyList.Push[per.Language] = langMsg
			notifyList.File[per.DeviceToken] = per.Language
			bodyData[per.Bundle] = notifyList
		} else {
			bodyData[per.Bundle] = &cst.NotifyT{
				Channel: per.Channel,
				Bundle:  per.Bundle,
				Push:    map[string]string{per.Language: langMsg},
				File:    map[string]string{per.DeviceToken: per.Language},
			}
		}
	}
	for _, dataCopy := range bodyData {
		//glog.Infof("dataCopy %v", dataCopy)
		_, err := restyHttp.Client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(dataCopy).
			Post(static.GetNotifyUrl())
		if err != nil {
			glog.Warnf("err %v", err)
			//直接返回
			return
		}
		//glog.Infof("rpc push return msg %s", string(postRes.Body()))
	}
	glog.Infof("send Ok tasks %d", len(tmpAllData))
}

// 添加进批量日志
func (t *taskCh) addTask(task *pbpush.PushUserTask) {
	t.allData = append(t.allData, task)
	if len(t.allData) >= msgNumsOk {
		t.doAllTasks()
	}
}
