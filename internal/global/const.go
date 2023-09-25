package global

const (
	// Mysql配置
	MysqlAddress  = "127.0.0.1"
	MysqlPort     = 3306
	MysqlUserName = "root"
	MysqlPassWord = "lin2738963"
	MysqlDataBase = "cloud-disk"

	// Redis配置
	RedisAddress = "127.0.0.1"
	RedisPort    = 6379

	// 服务器配置
	WebAddr = "127.0.0.1"
	WebPort = 8080

	// JWT鉴权配置
	JwtKey             = "cloud-disk-key"
	TokenExpire        = 3600
	RefreshTokenExpire = 7200

	Datetime = "2006-01-02 15:04:05"
)
