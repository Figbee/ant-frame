package initialize

import (
	"ant-frame/bootstrap/global"
	"fmt"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func InitCasbin() {
	var err error
	a, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		global.Logger.Panicf("权限模块载入失败:%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	global.Casbin, err = casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		global.Logger.Panicf("权限模块载入失败:%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	err = global.Casbin.LoadPolicy()
	if err != nil {
		global.Logger.Panicf("权限模块载入失败:%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	global.Logger.Info("权限模块载入成功")
}
