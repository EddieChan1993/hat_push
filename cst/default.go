package cst

import (
	"git.dhgames.cn/svr_comm/kite/utils/cast"
	"hat_push/pbgo/config"
	"strings"
)

const DefaultZone = "Asia/Shanghai" //默认时区
const DefPushClock = 7              //默认推送时间点
const (
	MMDBDebugPath     = "./tool/geolite2/GeoLite2-City.mmdb"
	MMDBDebugTestPath = "../tool/geolite2/GeoLite2-City.mmdb"
	MMDBPath          = "/home/dhcd/data/common/GeoIP2-City.mmdb"
)

const (
	ChIOS     = "ios"
	ChANDROID = "android"
	ChTPNS    = "tpns"
)

const (
	google  = "google"
	android = "android"
	ios     = "ios"
)

const (
	HandleType_OfflinePush = 1 //离线推送
)

func GetChannel(bound string) string {
	if strings.Contains(bound, google) {
		//海外android包
		return ChANDROID
	} else if strings.Contains(bound, ios) {
		return ChIOS
	} else {
		return ChTPNS
	}
}

// GetLangMsg 获取推送信息
func GetLangMsg(handleType int32, lang string) string {
	one := config.PushLanguageData.GetByHandleIdAndLang(handleType, cast.ToInt32(lang))
	if len(one) != 1 {
		return ""
	}
	return one[0].Msg
}
