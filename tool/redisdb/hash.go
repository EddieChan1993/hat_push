package redisdb

import (
	"git.dhgames.cn/svr_comm/gcore/gredis"
	"hat_push/tool/consul/static"
)

// HMSet 同时将多个 field-value
func HMSet(key string, fields map[string]interface{}) error {
	res, err := gredis.HMSet(static.StaticRedisUrl(), key, fields)
	if err != nil {
		return err
	}
	_, err = res.Result()
	if err != nil {
		return err
	}
	return nil
}

// HMDel 删除
func HMDel(key string, fields ...string) error {
	res, err := gredis.HDel(static.StaticRedisUrl(), key, fields...)
	if err != nil {
		return err
	}
	_, err = res.Result()
	if err != nil {
		return err
	}
	return nil
}

func HSet(key, field string, value interface{}) error {
	res, err := gredis.HSet(static.StaticRedisUrl(), key, field, value)
	if err != nil {
		return err
	}
	_, err = res.Result()
	if err != nil {
		return err
	}
	return nil
}

func HGet(key, field string) (string, error) {
	res, err := gredis.HGet(static.StaticRedisUrl(), key, field)
	if err != nil {
		return "", err
	}
	data, err := res.Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

// HGetAll 所有的字段和值
func HGetAll(key string) (map[string]string, error) {
	res, err := gredis.HGetAll(static.StaticRedisUrl(), key)
	if err != nil {
		return nil, err
	}
	resData, err := res.Result()
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// HIncrBy 增量
func HIncrBy(key, field string, nums int64) (int64, error) {
	res, err := gredis.HIncrBy(static.StaticRedisUrl(), key, field, nums)
	if err != nil {
		return 0, err
	}
	data, err := res.Result()
	if err != nil {
		return 0, err
	}
	return data, nil
}
