package util

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

func Logger() *zap.Logger {
	once.Do(func() {
		logger, _ = zap.NewProduction()
	})
	return logger
}
