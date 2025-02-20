package ants

import (
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	ants2 "github.com/panjf2000/ants/v2"
)

const poolSize = 100

var AntsGo *ants2.Pool

func InitAnts() {
	var err error
	AntsGo, err = ants2.NewPool(poolSize)
	if err != nil {
		klog.Panic(err)
	}
}

func ReleaseAnts() {
	AntsGo.Release()
}
