package log

var std Logger

func Debug(args ...interface{}) {
	std.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Info(args ...interface{}) {
	std.Info(args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Warn(args ...interface{}) {
	std.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Error(args ...interface{}) {
	std.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
func Panic(args ...interface{}) {
	std.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func InitGlobal(s Logger) {
	std = s
}
