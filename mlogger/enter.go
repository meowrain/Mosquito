package mlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var MLogger *zap.Logger

func InitLogger() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",                      // 时间字段的键名
		LevelKey:     "level",                     // 日志级别字段的键名
		MessageKey:   "msg",                       // 消息字段的键名
		CallerKey:    "caller",                    // 调用者字段的键名
		EncodeTime:   zapcore.ISO8601TimeEncoder,  // 时间格式
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 日志级别大写
		EncodeCaller: zapcore.ShortCallerEncoder,  // 调用者（文件名和行号）
	}
	// 创建 Console Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 将日志输出到标准输出
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.DebugLevel)
	MLogger = zap.New(core, zap.AddCaller())
	defer MLogger.Sync()
}
