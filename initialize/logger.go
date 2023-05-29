package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Info("初始化logger")
}