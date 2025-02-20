package util

import (
	"git.dhgames.cn/svr_comm/kite"
	"hat_push/tool/consul/dynamic"
)

// RpcTarWithRouting 路由表里会指定节点的Rpc目标
func RpcTarWithRouting(tar, from string, sid int64) *kite.Destination {
	dc, cl, _, _ := kite.GetWhoAmI()
	d := kite.Service(dc, cl, tar)
	n := dynamic.GetNode(tar, int(sid))
	if n != 0 {
		d.Index = n
	}
	return &d
}
