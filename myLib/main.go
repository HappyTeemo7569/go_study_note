package myLib

func init() {
	InitLog()
	Logger.Infof("日志初始化成功")

	InitConfig()
	Logger.Infof("配置初始化成功")

	InitRedis()
	Logger.Infof("redis初始化成功")

	InitMysqlNormal()
	Logger.Infof("mysql初始化成功")
}
