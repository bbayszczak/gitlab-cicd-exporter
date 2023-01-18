package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetLogLevel(logger *zap.SugaredLogger, atomLvl zap.AtomicLevel, lvl zapcore.Level) {
	logger.Infow("setting log level", "logLevel", lvl)
	atomLvl.SetLevel(lvl)
}
