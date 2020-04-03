package myLib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//获取RPC Redis连接
func GetRedisRpc() redis.Conn {
	rc := RedisPool.Get()
	rc.Do("SELECT", config.Redis.Rpc)
	return rc
}

//获取API Redis连接
func GetRedisApi() redis.Conn {
	rc := RedisPool.Get()
	rc.Do("SELECT", config.Redis.Api)
	return rc
}

func InitRedis() {
	//连接池
	RedisPool = redis.Pool{
		Dial: func() (redis.Conn, error) {
			connstr := fmt.Sprintf("%s:%d", config.Redis.Addr, config.Redis.Port)
			con, err := redis.Dial("tcp", connstr)
			if err != nil {
				//panic(err)
				return nil, err
			}
			if len(config.Redis.Pwd) > 0 {
				err = con.Send("auth", config.Redis.Pwd)
				if err != nil {
					//panic(err)
					return nil, err
				}
			}
			if _, err := con.Do("SELECT", config.Redis.Rpc); err != nil {
				//panic(err)
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         3000,
		MaxActive:       5000,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
