package myLib

import "fmt"

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMysqlNormal() {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.MySql.Auth, config.MySql.Pwd, config.MySql.Addr, config.MySql.Port, config.MySql.Db)

	var err error
	Db, err = sqlx.Open("mysql", connstr)
	if err != nil {
		panic(err)
		//Logger.Errorf("open mysql failed,", err)
		return
	}
}
