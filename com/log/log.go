package log

type CfgLog struct {
	FileName   string
	Level      string
	MaxSizeMB  int
	MaxFileNum int
	MaxFileDay int
	Compress   bool
	OnlyStdout bool // 开启后仅输出到 stdout
	Stdout     bool // 开启后同时输出到 stdout 和文件，否则仅输出到文件
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}
