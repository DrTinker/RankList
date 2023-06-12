package utils

import "github.com/go-redis/redis"

func RedisInit(addr, pwd string) *redis.Client {
	//初始化redis，连接地址和端口，密码，数据库名称
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})
	return rdb
}
