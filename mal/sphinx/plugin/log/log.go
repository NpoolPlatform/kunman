package log

import "github.com/NpoolPlatform/kunman/framework/logger"

func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}
