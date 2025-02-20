package config

type PushHandleCfg struct {
	Id        int32
	Parameter int32
}

type PushHandleTable struct {
	data map[int32]*PushHandleCfg
}

var PushHandleData = &PushHandleTable{
	data: map[int32]*PushHandleCfg{},
}

func (table *PushHandleTable) Get(id int32) *PushHandleCfg {
	return table.data[id]
}

func (table *PushHandleTable) GetAll() []int32 {
	return pushHandleKeys
}

func (table *PushHandleTable) GetAllData() []*PushHandleCfg {
	return pushHandleValues
}

var pushHandleValues = []*PushHandleCfg{
	{
		Id:        1,
		Parameter: 12,
	},
	{
		Id: 2,
	},
	{
		Id: 3,
	},
	{
		Id: 4,
	},
	{
		Id: 5,
	},
	{
		Id: 6,
	},
}

var pushHandleKeys = []int32{
	1, 2, 3, 4, 5, 6,
}

func init() {
	PushHandleData.data = make(map[int32]*PushHandleCfg)
	for i := 0; i < len(pushHandleKeys); i++ {
		PushHandleData.data[pushHandleKeys[i]] = pushHandleValues[i]
	}
}
