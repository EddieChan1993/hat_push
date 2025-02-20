package dynamic

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"git.dhgames.cn/svr_comm/kite"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
)

const (
	masterNodeFolder   = "masterNode"
	routingTableFolder = "routingTable"
)

var dynamicDirCfg = &DynamicDirCfg{}

type DynamicDirCfg struct {
	MasterNode   *MasterNode
	RoutingTable *RoutingTable
}

type MasterNode struct {
	Index int
}

type RoutingInfo map[string]map[int][]int

type RoutingTable struct {
	Info RoutingInfo `json:"info"`
}

func (p *DynamicDirCfg) Reload(key string, data []byte) {
	var err error
	switch key {
	case masterNodeFolder:
		masterNode := &MasterNode{}
		err = jsoniter.Unmarshal(data, masterNode)
		if err != nil {
			panic(err)
		}
		dynamicDirCfg.MasterNode = masterNode
	case routingTableFolder:
		routingTable := &RoutingTable{}
		err = jsoniter.Unmarshal(data, routingTable)
		if err != nil {
			panic(err)
		}
		dynamicDirCfg.RoutingTable = routingTable
	}
}

func (p *DynamicDirCfg) Delete(key string) {
}

func GetMasterNode() int {
	return dynamicDirCfg.MasterNode.Index
}

func GetDynamicDirCfg() *DynamicDirCfg {
	return dynamicDirCfg
}

// GetRoutingInfo 获取对应服务应该路由到哪个节点
func GetRoutingInfo() map[string]map[int]int {
	routingTableDataMap := map[string]map[int]int{}
	for key1, m1 := range dynamicDirCfg.RoutingTable.Info {
		if routingTableDataMap[key1] == nil {
			routingTableDataMap[key1] = map[int]int{}
		}
		for key2, val1 := range m1 {
			for _, sid := range val1 {
				routingTableDataMap[key1][sid] = key2
			}
		}
	}
	return routingTableDataMap
}

// GetNode 获取对应服务应该路由到哪个节点
func GetNode(service string, sid int) int {
	routingTableDataMap := map[string]map[int]int{}
	for key1, m1 := range dynamicDirCfg.RoutingTable.Info {
		if routingTableDataMap[key1] == nil {
			routingTableDataMap[key1] = map[int]int{}
		}
		for key2, val1 := range m1 {
			for _, sid := range val1 {
				routingTableDataMap[key1][sid] = key2
			}
		}
	}
	idx, ok := routingTableDataMap[service][sid]
	if !ok {
		return 1
	}
	return idx
}

//InitDynamicData 初始化值
func InitDynamicData(dir string, ser *consul.ServiceInfo) {
	res := GetDynamicDirCfg()
	//如果不存在，则主动写入
	if res.MasterNode == nil {
		data := &MasterNode{
			Index: ser.Index,
		}
		dir = dir + "/" + masterNodeFolder
		err := consul.UpdateDynamicConfigByDir(dir, data)
		if err != nil {
			klog.Panic(err)
		}
		res.MasterNode = data
	}
}

func AddToRoutingTable(sid int) {
	// 找出对应存活的节点，把sid放到最大的那个节点里
	routingMap := GetRoutingInfo()
	if routingMap != nil && routingMap["re_guild"][sid] != 0 {
		return
	}
	// 找到目标服务的最大节点
	dc, cl, _, _ := kite.GetWhoAmI()
	alive := kite.Alive(kite.Service(dc, cl, "re_guild"))
	if len(alive) < 1 {
		return
	}
	maxIdx := -1
	for _, node := range alive {
		s := strings.Split(node, "-")
		idx, _ := strconv.Atoi(s[2])
		if idx > maxIdx {
			maxIdx = idx
		}
	}
	if maxIdx == -1 {
		return
	}
	err := JustAddRoutingKey("re_guild", maxIdx, []int{sid})
	if err != nil {
		return
	}
	return
}

func JustAddRoutingKey(svrName string, idx int, newKeys []int) error {
	newInfo := map[int][]int{}
	for _, key := range newKeys {
		err := nodeAdd(svrName, newInfo, idx, key)
		if err != nil {
			klog.Errorf("添加节点时出错了 svr %s, idx %s, key %s, err %s", svrName, idx, key, err)
			return err
		}
	}
	return UpdateTable(svrName, idx, newInfo)
}

func nodeAdd(svr string, newInfo map[int][]int, to int, key int) error {
	// 如果没有，则取现有的
	l, ok := newInfo[to]
	if !ok {
		newInfo[to] = dynamicDirCfg.RoutingTable.Info[svr][to]
		l = newInfo[to]
	}
	for _, v := range l {
		if v == key {
			return nil
		}
	}
	l = append(l, key)
	newInfo[to] = l
	return nil
}

// UpdateTable 把新数据写到consul
func UpdateTable(key1 string, idx int, v map[int][]int) error {
	if dynamicDirCfg.RoutingTable.Info == nil {
		dynamicDirCfg.RoutingTable.Info = map[string]map[int][]int{}
	}
	dynamicDirCfg.RoutingTable.Info[key1][idx] = v[idx]
	_, cl, ser, _ := kite.GetWhoAmI()
	dynamicDir := cl + "/" + ser + "/" + routingTableFolder
	return consul.UpdateDynamicConfigByDir(dynamicDir, dynamicDirCfg.RoutingTable)
}
