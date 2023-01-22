package logging

import (
	"os"

	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (*zap.Logger, zap.AtomicLevel) {
	atom := zap.NewAtomicLevel()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))
	return logger, atom
}

func SetLogLevel(logger *zap.SugaredLogger, atomLvl zap.AtomicLevel, lvl zapcore.Level) {
	logger.Infow("setting log level", "logLevel", lvl)
	atomLvl.SetLevel(lvl)
}

func getIDFromContext(c echo.Context) string {
	id := c.Request().Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = c.Response().Header().Get(echo.HeaderXRequestID)
	}
	return id
}

func GetLoggerFromContext(c echo.Context) *zap.SugaredLogger {
	cc := c.(*customcontext.CustomContext)
	return cc.Log.With("id", getIDFromContext(c))
}
