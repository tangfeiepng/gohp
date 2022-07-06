package logger

import (
	"Walker/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

//日志接口

//对接口进行实现
var logger *zap.Logger

func InitLogger() *zap.Logger {
	core := zapcore.NewCore(getEncoder(), getWriterLogger(), zap.DebugLevel)
	logger = zap.New(core)
	return logger
}

//修改zap原有的写入器已方便支持日志分割
func getWriterLogger() zapcore.WriteSyncer {
	date := time.Now().Format("2006-01-02")
	writerLogger := lumberjack.Logger{
		Filename: global.StoragePath + "t_log_" + date + ".log",
	}
	return zapcore.AddSync(&writerLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "date"
	//输出json格式日志
	return zapcore.NewJSONEncoder(encoderConfig)
}
