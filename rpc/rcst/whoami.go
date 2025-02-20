package rcst

import (
	"git.dhgames.cn/svr_comm/kite"
)

const (
	GateSvrPush = "hat_push"
)

func RpcPushWho() *kite.Destination {
	return whoAmISvrNameIndex(GateSvrPush)
}

func whoAmISvrNameIndex(service string, index ...int) *kite.Destination {
	dc, cl, _, _ := kite.GetWhoAmI()
	des := &kite.Destination{
		DC:      dc,
		Cluster: cl,
		Service: service,
	}
	if len(index) == 0 {
		return des
	}
	des.Index = index[0]
	return des
}
