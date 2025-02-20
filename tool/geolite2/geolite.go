package geolite2

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gcore/glog"
	"github.com/oschwald/geoip2-golang"
	"hat_push/cst"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var geoip2Db *geoip2.Reader
var geoIpPool sync.Map

// InitGeoIp ip时区转化
func InitGeoIp() {
	argsLen := len(os.Args)
	if argsLen > 3 && os.Args[argsLen-1] == "debug" {
		geoipIns(cst.MMDBDebugPath)
	} else if argsLen > 3 && strings.Contains(os.Args[argsLen-2], "-test") {
		geoipIns(cst.MMDBDebugTestPath)
	} else {
		geoipIns(cst.MMDBPath)
	}
}

func GetTimeZone(ip string) (string, error) {
	ipP := net.ParseIP(ip)
	record, err := geoip2Db.City(ipP)
	if err != nil {
		return "", err
	}
	if record.Location.TimeZone == "" {
		return cst.DefaultZone, nil
	}
	return record.Location.TimeZone, nil
}

// ThinkSendAt 处理发送时间到合理的时间点发送
func ThinkSendAt(sendCd int64, ip string) (int64, error) {
	zoneTime, err := GetTimeZone(ip)
	if err != nil {
		glog.Error(ip, err)
		return 0, err
	}
	cstSh, had := getGeoIpPool(zoneTime)
	if !had {
		return 0, fmt.Errorf("getGeoIpPool none ip %s", ip)
	}
	timeT := time.Now().In(cstSh)
	sendAt := timeT.Unix() + sendCd //当前ip下的发送时间
	sendTimeT := time.Unix(sendAt, 0).In(cstSh)
	h, _, _ := sendTimeT.Clock()
	if h >= 1 && h <= 6 {
		//1点～6点，调整到7点推送
		tm1 := time.Date(sendTimeT.Year(), sendTimeT.Month(), sendTimeT.Day(), cst.DefPushClock, 0, 0, 0, timeT.Location())
		return tm1.Unix(), nil
	} else {
		return sendAt, nil
	}
}

func CloseGeoIp() {
	err := geoip2Db.Close()
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Info("geoip2Db关闭")
}

func geoipIns(dbPath string) {
	var err error
	geoip2Db, err = geoip2.Open(dbPath)
	if err != nil {
		glog.Panic(err)
	}
}

func getGeoIpPool(timeZone string) (*time.Location, bool) {
	one, had := geoIpPool.Load(timeZone)
	if had {
		return one.(*time.Location), true
	} else {
		cstSh, err := time.LoadLocation(timeZone)
		if err != nil {
			//glog.Error(err)
			return nil, false
		}
		geoIpPool.Store(timeZone, cstSh)
		return cstSh, true
	}
}
