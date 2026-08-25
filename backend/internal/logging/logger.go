package logging

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger() {
	var config zap.Config

	env := os.Getenv("ENV")
	if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}
	}

	logger, err := config.Build()
	if err != nil {
		// Fallback to a basic logger instead of panicking
		fallback := zap.NewNop().Sugar()
		Log = fallback
		Log.Warnf("Failed to initialize zap logger, using noop logger: %v", err)
		return
	}

	Log = logger.Sugar()
}

func Infof(template string, args ...interface{}) {
	Log.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Log.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Log.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Log.Fatalf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	Log.Debugf(template, args...)
}

// RouteLogger handles printing custom route log attributes in HTTP middleware
func RouteLogger(c echo.Context, values MiddlewareLogValues) error {
	Log.Infow("HTTP Request",
		"method", values.Method,
		"uri", values.URI,
		"status", values.Status,
		"latency", values.Latency,
		"ip", values.IP,
		"error", values.Error,
	)
	return nil
}

type MiddlewareLogValues struct {
	Method  string
	URI     string
	Status  int
	Latency time.Duration
	IP      string
	Error   error
}
