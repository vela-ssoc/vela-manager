package conf

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Level     string `json:"level"     yaml:"level"`
	Console   bool   `json:"console"   yaml:"console"`
	Colorful  bool   `json:"colorful"  yaml:"colorful"`
	Directory string `json:"directory" yaml:"directory"`
	Maxsize   int    `json:"maxsize"   yaml:"maxsize"`
	MaxAge    int    `json:"maxage"    yaml:"maxage"`
	Backup    int    `json:"backup"    yaml:"backup"`
	Localtime bool   `json:"localtime" yaml:"localtime"`
	Compress  bool   `json:"compress"  yaml:"compress"`
}

func (l Logger) Zap() *zap.Logger {
	console := l.Console
	var filename string
	if dir := l.Directory; dir != "" {
		filename = filepath.Join(dir, "manager.log")
	}
	// 既不输出到控制台又不输出到日志文件
	if !console && filename == "" {
		return zap.NewNop()
	}

	prod := zap.NewProductionEncoderConfig()
	prod.EncodeTime = zapcore.ISO8601TimeEncoder
	if l.Colorful {
		prod.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		prod.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	var syncer zapcore.WriteSyncer
	if console {
		syncer = zapcore.AddSync(os.Stdout)
	}
	if filename != "" {
		lumber := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    l.Maxsize,
			MaxAge:     l.MaxAge,
			MaxBackups: l.Backup,
			LocalTime:  l.Localtime,
			Compress:   l.Compress,
		}
		if syncer == nil {
			syncer = zapcore.AddSync(lumber)
		} else {
			syncer = zapcore.NewMultiWriteSyncer(syncer, zapcore.AddSync(lumber))
		}
	}

	encoder := zapcore.NewConsoleEncoder(prod)
	level := zapcore.WarnLevel
	_ = level.Set(l.Level) // 就算设置失败还是默认值 WarnLevel
	core := zapcore.NewCore(encoder, syncer, level)

	return zap.New(core, zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel))
}
