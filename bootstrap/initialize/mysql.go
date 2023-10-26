package initialize

import (
	"ant-frame/bootstrap/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlConf = global.Conf.Mysql
)

func InitMysql() {
	dsn := dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Fatalf("mysql connect error: %v", err)
		panic("数据库连接失败")
	}
	//开启调试模式
	if mysqlConf.LogMode {
		db = db.Debug()
	}
	global.DB = db

}

/*
mysql 连接字符串
*/
func dsn() string {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DbName,
		mysqlConf.Query)
	return dataSource
}
