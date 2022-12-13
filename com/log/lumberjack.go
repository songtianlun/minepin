package log

import "github.com/natefinch/lumberjack"

func getLumberJackLogger(cfg *CfgLog) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxFileNum,
		MaxAge:     cfg.MaxFileDay,
		Compress:   cfg.Compress,
	}
}
