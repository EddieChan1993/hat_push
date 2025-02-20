package comDb

import "hat_push/util"

func collection(table string) string {
	return util.LogicName + "-" + table
}

// InitComDb 公共数据存储
func InitComDb() {
	initDbWorker()
	initPushTaskDb()
}
