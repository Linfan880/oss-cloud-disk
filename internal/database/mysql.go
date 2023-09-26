package database

import (
	"file-system/internal/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		global.MysqlUserName,
		global.MysqlPassWord,
		global.MysqlAddress,
		global.MysqlPort,
		global.MysqlDataBase)

	var err error
	global.MysqlConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Mysql Connection Failed ...")
		return
	}
	fmt.Println("Mysql Connection Success ...")
}

func GetDbConn() *gorm.DB {
	return global.MysqlConn
}
