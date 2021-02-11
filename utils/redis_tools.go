package utils

import (
	"blog/sysinit"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

// 初始化redis连接池
var rdb = sysinit.RedisPool()

/**
redis  SET
*/
func RdbSet(key, v string) (string, error) {
	b, err := redis.String(rdb.Do("SET", key, v))
	if err != nil {
		logs.Error("set error", err.Error())
		return "", err
	}
	return b, nil
}

/**
redis SET
设置键值并给出过期时间
*/
func RdbSetExp(key, v string, ex int) error {
	_, err := rdb.Do("SET", key, v, "EX", ex)
	if err != nil {
		logs.Error("set error", err.Error())
		return err
	}
	return nil
}

/**
redis  GET
*/
func RdbGet(key string) (string, error) {
	val, err := redis.String(rdb.Do("GET", key))
	if err != nil {
		logs.Error("get error", err.Error())
		return "", err
	}
	return val, nil
}

/**
redis EXPIRE
给key设置过期时间
*/
func RdbSetKeyExp(key string, ex int) error {
	_, err := rdb.Do("EXPIRE", key, ex)
	if err != nil {
		logs.Error("set error", err.Error())
		return err
	}
	return nil
}

/**
redis  DEL
删除key
*/
func RdbDel(key string) (int64, error) {
	val, err := redis.Int64(rdb.Do("Del", key))
	if err != nil {
		logs.Error("del error", err.Error())
		return 0, err
	}
	return val, nil
}

/**
redis  EXISTS
检查key是否存在值
*/
func RdbCheck(key string) (bool, error) {
	b, err := redis.Bool(rdb.Do("EXISTS", key))
	if err != nil {
		logs.Error("EXISTS error", err.Error())
		return false, err
	}
	return b, nil
}

/**
redis LPush
将value插入队列头部
*/
func RdbLPush(key, v string) (int64, error) {
	val, err := redis.Int64(rdb.Do("LPUSH", key, v))
	if err != nil {
		logs.Error("LPUSH error", err.Error())
		return 0, err
	}
	return val, nil
}

/**
redis LPop
移出并获取队列的第一个元素
*/
func RdbLPop(key, v string) (int64, error) {
	val, err := redis.Int64(rdb.Do("LPOP", key))
	if err != nil {
		logs.Error("LPOP error", err.Error())
		return 0, err
	}
	return val, nil
}

/**
redis LLen
获取队列长度
*/
func RdbLLen(key, v string) (int64, error) {
	val, err := redis.Int64(rdb.Do("LLEN", key))
	if err != nil {
		logs.Error("LLEN error", err.Error())
		return 0, err
	}
	return val, nil
}
