package global

import (
	"ant-frame/bootstrap/config"
	"go.uber.org/zap"
)

var (
	Conf   config.Config
	Logger *zap.SugaredLogger
)
