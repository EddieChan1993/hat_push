package core

import (
	"context"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"hat_push/cst"
	"hat_push/pbgo/pbpush"
	"hat_push/tool/comDb"
	"hat_push/tool/geolite2"
	"hat_push/tool/goRuntime"
	"hat_push/util"
	"time"
)

/*
*
任务分派
*/
const tickTime = 6
const hour12 = 12 * 60 * 60

func initDispatch() {
	goRuntime.GoRun(func(ctx context.Context) {
		tick := time.NewTicker(tickTime * time.Second)
		defer tick.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				dispatch()
			}
		}
	})
}

func dispatch() {
	tasks := comDb.FindTasksBeforeAt(util.TimeNowUnix())
	if len(tasks) == 0 {
		return
	}
	var m int32
	var worker *taskCh
	for _, task := range tasks {
		//随机分发
		m = util.RandInt32(100) % workersNums
		worker = taskWorkers[m]
		worker.taskCh <- task
		oldTaskThink(task)
	}
}

func oldTaskThink(task *pbpush.PushUserTask) {
	var err error
	sendAt := task.SendAt
	if task.HandleType == cst.HandleType_OfflinePush {
		//间隔12个小时,继续推送
		task.SendAt, err = geolite2.ThinkSendAt(hour12, task.Ip)
		if err != nil {
			//klog.Warn(err)
			return
		}
		err = comDb.UpdatePushTask(comDb.PushTaskIsSameFileter(task.Account, task.HandleType, sendAt), task)
		if err != nil {
			klog.Warn(err)
			return
		}
	} else {
		//删掉，为了避免玩家刚好刷新了新的下轮push数据，此处将sendAt作为筛选条件，用于确定数据是否已经被修改，则不删除
		comDb.DelPushTask(comDb.PushTaskIsSameFileter(task.Account, task.HandleType, sendAt))
		return
	}
}
