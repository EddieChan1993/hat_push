package api

import (
	"git.dhgames.cn/svr_comm/kite"
	"hat_push/cst"
	"hat_push/pbgo/pbpush"
	"hat_push/tool/comDb"
	"hat_push/tool/geolite2"
)

type PushApi struct {
	meta *kite.Meta
}

func (p *PushApi) New() pbpush.PushApiServer {
	return &PushApi{}
}

func (p *PushApi) SetMeta(meta *kite.Meta) {
	p.meta = meta
}

func (p *PushApi) AddPushTasks(req *pbpush.ReqAddPushTasks) (*pbpush.RspAddPushTasks, error) {
	res := &pbpush.RspAddPushTasks{
		Status: 0,
		ErrMsg: "",
	}
	var sendAt int64
	var err error
	for _, data := range req.Tasks {
		if data.Account == 0 || data.DeviceToken == "" || data.Bundle == "" || data.SendCd <= 0 {
			continue
		}
		sendAt, err = geolite2.ThinkSendAt(data.SendCd, data.Ip)
		if err != nil {
			//klog.Warnf("%d ip %s sendCd %d err %v", data.Account, data.Ip, data.SendCd, err)
			continue
		}
		if sendAt <= 0 {
			//klog.Warnf("%d ip %s sendCd %d err %v", data.Account, data.Ip, data.SendCd, err)
			continue
		}
		data.Channel = cst.GetChannel(data.Bundle)
		data.SendAt = sendAt
		comDb.PushWorkerToSaveDb(data)
	}
	return res, nil
}

func (p *PushApi) DelPushTasks(req *pbpush.ReqDelPushTaskAll) (*pbpush.RspDelPushTaskAll, error) {
	comDb.DelPushTask(comDb.PushTaskAccount(req.Account))
	return nil, nil
}

func (p *PushApi) DelPushTaskType(req *pbpush.ReqDelPushTaskType) (*pbpush.RspDelPushTaskType, error) {
	comDb.DelPushTask(comDb.PushTaskFileter(req.Account, req.HandleType))
	return nil, nil
}
