package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"minepin/com/cfg"
	"os"
)

type logCfg struct {
	fileName   string
	level      string
	maxSizeMb  int
	maxFileNum int
	maxFileDay int
	compress   bool
}

var lCfg *logCfg

var Logger *zap.Logger             // 在每一微秒和每一次内存分配都很重要的上下文中，使用Logger，只支持强类型的结构化日志记录
var SugarLogger *zap.SugaredLogger // 在性能很好但不是很关键的上下文中，使用SugaredLogger，支持结构化和 printf 风格
var logLevel zapcore.LevelEnabler

func InitLogger(n string, l string, msm int, mfn int, mfd int, c bool) {
	lCfg = &logCfg{
		fileName:   n,
		level:      l,
		maxSizeMb:  msm,
		maxFileNum: mfn,
		maxFileDay: mfd,
		compress:   c,
	}

	switch lCfg.level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	core := zapcore.NewTee(getAllCores()...)

	// AddCaller - 调用函数信息记录到日志中的功能
	// AddCallerSkip - 向上跳 1 层，输出调用封装日志函数的位置
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	SugarLogger = Logger.Sugar()
	defer Logger.Sync() // flushes buffer, if any
	defer SugarLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 格式化时间显示
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   lCfg.fileName,
		MaxSize:    lCfg.maxSizeMb,
		MaxBackups: lCfg.maxFileNum,
		MaxAge:     lCfg.maxFileDay,
		Compress:   lCfg.compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getAllCores() []zapcore.Core {
	var allCore []zapcore.Core

	encoder := getEncoder()
	// debug 模式、显式输出到stdout 或 仅输出到 stdout 时将日志同时输出到 stdout
	if cfg.GetString("log.level") == "debug" ||
		cfg.GetBool("log.stdout") ||
		cfg.GetBool("log.only_stdout") {
		consoleDebugging := zapcore.Lock(os.Stdout)
		allCore = append(allCore, zapcore.NewCore(encoder, consoleDebugging, logLevel))
	}
	// 仅输出到 stdout 时屏蔽文件输入
	if !cfg.GetBool("log.only_stdout") {
		writeSyncer := getLogWriter()
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, logLevel))
	}
	return allCore
}

func DebugF(format string, v ...interface{}) { SugarLogger.Debugf(format, v...) }
func InfoF(format string, v ...interface{})  { SugarLogger.Infof(format, v...) }
func WarnF(format string, v ...interface{})  { SugarLogger.Warnf(format, v...) }
func ErrorF(format string, v ...interface{}) { SugarLogger.Errorf(format, v...) }
func PanicF(format string, v ...interface{}) { SugarLogger.Panicf(format, v...) }
func FatalF(format string, v ...interface{}) { SugarLogger.Fatalf(format, v...) }
func Debug(v ...interface{})                 { SugarLogger.Debug(v...) }
func Info(v ...interface{})                  { SugarLogger.Info(v...) }
func Warn(v ...interface{})                  { SugarLogger.Warn(v...) }
func Error(v ...interface{})                 { SugarLogger.Error(v...) }
func Panic(v ...interface{})                 { SugarLogger.Panic(v...) }
func Fatal(v ...interface{})                 { SugarLogger.Fatal(v...) }
