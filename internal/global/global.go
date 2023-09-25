package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var MysqlConn *gorm.DB
var RedisConn *redis.Client
