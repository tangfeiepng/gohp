package logger

import (
	"Walker/pkg/contract"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapClass struct {
	config contract.Config
	app    contract.Application
	logger *zap.Logger
}

func (zapClass *ZapClass) InitLogger() contract.Logger {
	//编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "created_at"
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}
	//写入器
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   zapClass.app.Get("path.storage").(string) + "/logs/" + zapClass.config.GetString("logger.filename"),
		MaxSize:    zapClass.config.GetInt("logger.max_size"),   //单个日志文件的大小
		MaxAge:     zapClass.config.GetInt("logger.max_age"),    //文件保存天数
		MaxBackups: zapClass.config.GetInt("logger.max_backup"), //文件保存个数
		Compress:   zapClass.config.GetBool("logger.compress"),  //文件是否压缩
	})
	zapClass.logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, zap.InfoLevel))
	return zapClass
}
func (zapClass *ZapClass) Debug(msg string, inf ...contract.LoggerField) {
	//把额外字段转成需要得类型
	fields := zapClass.conversionField(inf...)
	zapClass.logger.Debug(msg, fields...)
}
func (zapClass *ZapClass) Info(msg string, inf ...contract.LoggerField) {
	fields := zapClass.conversionField(inf...)
	zapClass.logger.Info(msg, fields...)
}
func (zapClass *ZapClass) Warn(msg string, inf ...contract.LoggerField) {
	fields := zapClass.conversionField(inf...)
	zapClass.logger.Warn(msg, fields...)
}
func (zapClass *ZapClass) Error(msg string, inf ...contract.LoggerField) {
	fields := zapClass.conversionField(inf...)
	zapClass.logger.Error(msg, fields...)
}

func (zapClass *ZapClass) conversionField(fields ...contract.LoggerField) []zap.Field {
	zapField := make([]zap.Field, 0)
	for _, field := range fields {
		//进行断言
		zapField = append(zapField, zap.Reflect(field.Key, field.Val))
	}
	return zapField
}
