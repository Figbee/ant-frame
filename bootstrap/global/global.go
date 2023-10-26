package global

import (
	"ant-frame/bootstrap/config"

	"github.com/casbin/casbin/v2"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	Conf   config.Config
	Logger *zap.SugaredLogger
	DB     *gorm.DB
	Casbin *casbin.Enforcer
)
