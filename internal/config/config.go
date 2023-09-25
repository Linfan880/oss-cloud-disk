package config

type Config struct {
	Mysql MysqlConfig
	Redis RedisConfig
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
