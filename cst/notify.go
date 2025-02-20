package cst

/**
推送消息结构体
*/

type NotifyT struct {
	Channel string            `json:"channel"`
	Bundle  string            `json:"bundle"`
	Push    map[string]string `json:"push"`
	File    map[string]string `json:"file"`
}

type Reply struct {
	SuccessCnt int           `json:"success_count"`
	FailCnt    int           `json:"fail_count"`
	Fail       []FailMessage `json:"fails"`
}

type FailMessage struct {
	Id      string `json:"id"`
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}
