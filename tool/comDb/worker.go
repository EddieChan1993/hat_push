package comDb

/**
数据存储缓冲池
*/
import (
	"context"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"hat_push/pbgo/pbpush"
	"hat_push/tool/goRuntime"
)

var dbWork *dbWorker

const maxBuffer = 800

type dbWorker struct {
	pushTasksAddCh chan *pbpush.PushUserTask //推送任务数据存储
}

func initDbWorker() {
	klog.Infof("Tool initDbWorker Run")
	dbWork = &dbWorker{
		pushTasksAddCh: make(chan *pbpush.PushUserTask, maxBuffer),
	}
	dbWork.goRun()
	klog.Infof("Tool initDbWorker Run Ok")
}

func (this_ *dbWorker) goRun() {
	goRuntime.GoRun(func(ctx context.Context) {
		defer func() {
			this_.releaseWorker()
			klog.Infof("Tool InitDbWorker Close")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-this_.pushTasksAddCh:
				if data == nil {
					continue
				}
				err := SetTopDigestBotCamp(data)
				if err != nil {
					klog.Warnf("SetDigestCamp err %v", err)
				}
			}
		}
	})
}

// releaseWorker 释放worker数据
func (this_ *dbWorker) releaseWorker() {
}

// PushWorkerToSaveDb 阵营数据存入
func PushWorkerToSaveDb(data *pbpush.PushUserTask) {
	dbWork.pushTasksAddCh <- data
}
