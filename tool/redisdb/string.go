package redisdb

import (
	"git.dhgames.cn/svr_comm/gcore/gredis"
	"hat_push/tool/consul/static"
)

func Set(key, val string) error {
	res, err := gredis.Set(static.StaticRedisUrl(), key, val, 0)
	if err != nil {
		return err
	}
	_, err = res.Result()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	res, err := gredis.Get(static.StaticRedisUrl(), key)
	if err != nil {
		return "", err
	}
	resData, err := res.Result()
	if err != nil {
		return "", err
	}
	return resData, nil
}
