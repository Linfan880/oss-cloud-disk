package config

type Config struct {
	Mysql MysqlConfig
	Redis RedisConfig
	Web   WebConfig
}

type MysqlConfig struct {
	Address  string
	Port     int
	UserName string
	PassWord string
	DataBase string
}

type RedisConfig struct {
	Address string
	Port    int
}

// 服务器配置
type WebConfig struct {
	Address string
	Port    int
}
