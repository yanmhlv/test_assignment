package app

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once    sync.Once
	sugared *zap.SugaredLogger
)

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		sugared = logger.Sugar()
	})
	return sugared
}
