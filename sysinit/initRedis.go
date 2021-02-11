package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

/**
redis 初始化连接
*/
func getRedisPool() *redis.Pool {
	redisHost := beego.AppConfig.String("redis::redis_host")
	redisPort := beego.AppConfig.String("redis::redis_port")
	redisPassword := beego.AppConfig.String("redis::redis_pwd")

	//建立连接池
	return &redis.Pool{
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: beego.AppConfig.DefaultInt("redis::redis_max_idle", 1),
		//最大的激活连接数，表示同时最多有N个连接
		MaxActive: beego.AppConfig.DefaultInt("redis::redis_max_active", 10),
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 300 * time.Second,
		//建立连接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost+":"+redisPort)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			if redisPassword != "" {
				if _, authErr := c.Do("AUTH", redisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			return c, nil
		},
	}
}

// 对外暴露线连接池
func RedisPool() redis.Conn {
	return getRedisPool().Get()
}
