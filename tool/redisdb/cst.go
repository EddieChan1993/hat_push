package redisdb

import (
	"fmt"
	"hat_push/tool/consul"
)

const TimeMaxLen = 10000000000 //时间戳最大长度

type ZSetScoreTyp = float64

// GetRedisKey 当前集群当前服务
func GetRedisKey(key string, sid int32) string {
	serviceInfo := consul.WhoAmI()
	return fmt.Sprintf("%s:%s:%d:%s", serviceInfo.Cluster, serviceInfo.Service, sid, key)
}

// GetRedisKeyByZone 赛区关联
func GetRedisKeyByZone(key string, zoneId int32) string {
	serviceInfo := consul.WhoAmI()
	return fmt.Sprintf("%s:%s:zoneId:%d:%s", serviceInfo.Cluster, serviceInfo.Service, zoneId, key)
}

// GetRedisKeyGlobal 全服key
func GetRedisKeyGlobal(key string) string {
	serviceInfo := consul.WhoAmI()
	return fmt.Sprintf("%s:%s:%s", serviceInfo.Cluster, serviceInfo.Service, key)
}
