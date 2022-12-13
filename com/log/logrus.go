package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Fields map[string]interface{}

type logrusAdapt struct {
	l    *logrus.Logger
	lCfg *CfgLog
}

func (l *logrusAdapt) WithField(key string, value interface{}) Logger {
	return newFieldAdapt(l.l.WithField(key, value))
}

func (l *logrusAdapt) Debugf(format string, v ...interface{}) { l.l.Debugf(format, v...) }
func (l *logrusAdapt) Infof(format string, v ...interface{})  { l.l.Infof(format, v...) }
func (l *logrusAdapt) Warnf(format string, v ...interface{})  { l.l.Warnf(format, v...) }
func (l *logrusAdapt) Errorf(format string, v ...interface{}) { l.l.Errorf(format, v...) }
func (l *logrusAdapt) Panicf(format string, v ...interface{}) { l.l.Panicf(format, v...) }
func (l *logrusAdapt) Fatalf(format string, v ...interface{}) { l.l.Fatalf(format, v...) }
func (l *logrusAdapt) Debug(v ...interface{})                 { l.l.Debug(v...) }
func (l *logrusAdapt) Info(v ...interface{})                  { l.l.Info(v...) }
func (l *logrusAdapt) Warn(v ...interface{})                  { l.l.Warn(v...) }
func (l *logrusAdapt) Error(v ...interface{})                 { l.l.Error(v...) }
func (l *logrusAdapt) Panic(v ...interface{})                 { l.l.Panic(v...) }
func (l *logrusAdapt) Fatal(v ...interface{})                 { l.l.Fatal(v...) }

// 封装logrus.Entry
type fieldAdapt struct {
	e *logrus.Entry
}

func (f fieldAdapt) WithField(key string, value interface{}) Logger {
	return newFieldAdapt(f.e.WithField(key, value))
}

func (f fieldAdapt) WithFields(fields Fields) Logger {
	return newFieldAdapt(f.e.WithFields(logrus.Fields(fields)))
}

//func (f fieldAdapt) Tracef(format string, args ...interface{}) {
//    panic("implement me")
//}

func (f fieldAdapt) WithError(err error) Logger {
	return newFieldAdapt(f.e.WithError(err))
}

func (f fieldAdapt) Debugf(format string, args ...interface{}) {
	f.e.Debugf(format, args...)
}

func (f fieldAdapt) Infof(format string, args ...interface{}) {
	f.e.Infof(format, args...)
}

func (f fieldAdapt) Printf(format string, args ...interface{}) {
	f.e.Printf(format, args...)
}

func (f fieldAdapt) Warnf(format string, args ...interface{}) {
	f.e.Warnf(format, args...)
}

func (f fieldAdapt) Warningf(format string, args ...interface{}) {
	f.e.Warningf(format, args...)
}

func (f fieldAdapt) Errorf(format string, args ...interface{}) {
	f.e.Errorf(format, args...)
}

func (f fieldAdapt) Fatalf(format string, args ...interface{}) {
	f.e.Fatalf(format, args...)
}

func (f fieldAdapt) Panicf(format string, args ...interface{}) {
	f.e.Panicf(format, args...)
}

func (f fieldAdapt) Debug(args ...interface{}) {
	f.e.Debug(args...)
}

func (f fieldAdapt) Info(args ...interface{}) {
	f.e.Info(args...)
}

func (f fieldAdapt) Print(args ...interface{}) {
	f.e.Print(args...)
}

func (f fieldAdapt) Warn(args ...interface{}) {
	f.e.Warn(args...)
}

func (f fieldAdapt) Warning(args ...interface{}) {
	f.e.Warning(args...)
}

func (f fieldAdapt) Error(args ...interface{}) {
	f.e.Error(args...)
}

func (f fieldAdapt) Fatal(args ...interface{}) {
	f.e.Fatal(args...)
}

func (f fieldAdapt) Panic(args ...interface{}) {
	f.e.Panic(args...)
}

func (f fieldAdapt) Debugln(args ...interface{}) {
	f.e.Debugln(args...)
}

func (f fieldAdapt) Infoln(args ...interface{}) {
	f.e.Infoln(args...)
}

func (f fieldAdapt) Println(args ...interface{}) {
	f.e.Println(args...)
}

func (f fieldAdapt) Warnln(args ...interface{}) {
	f.e.Warnln(args...)
}

func (f fieldAdapt) Warningln(args ...interface{}) {
	f.e.Warningln(args...)
}

func (f fieldAdapt) Errorln(args ...interface{}) {
	f.e.Errorln(args...)
}

func (f fieldAdapt) Fatalln(args ...interface{}) {
	f.e.Fatalln(args...)
}

func (f fieldAdapt) Panicln(args ...interface{}) {
	f.e.Panicln(args...)
}

func (f fieldAdapt) Trace(args ...interface{}) {
	f.e.Trace(args...)
}

func newFieldAdapt(e *logrus.Entry) Logger {
	return fieldAdapt{e}
}

func newLogrus(lCfg *CfgLog) *logrus.Logger {
	l := logrus.New()

	// Step1 设置日志级别
	level, err := logrus.ParseLevel(lCfg.Level)
	if err != nil {
		l.SetLevel(logrus.DebugLevel)
	}
	l.SetLevel(level)

	// Step2 设置日志格式
	showPosition := false
	if level == logrus.DebugLevel {
		showPosition = true
	}
	l.SetFormatter(&Formatter{
		FieldsOrder:           nil,
		TimeStampFormat:       "2006-01-02 15:04:05",
		CharStampFormat:       "",
		HideKeys:              false,
		Position:              showPosition,
		PositionSkip:          11, // 直接调用跳过5层，接口封装跳过6层
		Colors:                true,
		FieldsColors:          true,
		FieldsSpace:           true,
		ShowFullLevel:         false,
		LowerCaseLevel:        false,
		TrimMessages:          true,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
	})

	// Step3 设置日志输出
	if lCfg.OnlyStdout { // 仅输出到 stdout
		l.SetOutput(os.Stdout)
	} else if lCfg.Stdout { // 输出到 stdout 和文件
		l.SetOutput(io.MultiWriter(os.Stdout, getLumberJackLogger(lCfg)))
	} else { // 仅输出到文件
		l.SetOutput(getLumberJackLogger(lCfg))
	}

	//l.SetReportCaller(true) // 开启后会显示详细的位置

	return l
}

func NewLogrus(lCfg *CfgLog) Logger {
	return &logrusAdapt{
		l:    newLogrus(lCfg),
		lCfg: lCfg,
	}
}
