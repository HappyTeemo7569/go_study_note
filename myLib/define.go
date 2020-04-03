package myLib

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

//数据库连接
var Db *sqlx.DB

//Redis连接池
var RedisPool redis.Pool

var config tagConfig

type tagConfig struct {
	MySql tagMysqlConfig
	Redis tagRedisConfig
	Other tagOtherConfig
}

type tagMysqlConfig struct {
	Auth string
	Pwd  string
	Addr string
	Port int
	Db   string
}

type tagRedisConfig struct {
	Addr string
	Port int
	Pwd  string
	Rpc  int
	Api  int
}

type tagOtherConfig struct {
	Debug bool
}
