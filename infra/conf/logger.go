package conf

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Level    string             `json:"level"    yaml:"level"`
	Console  bool               `json:"console"  yaml:"console"`
	Colorful bool               `json:"colorful" yaml:"colorful"`
	Lumber   *lumberjack.Logger `json:"lumber"   yaml:"lumber"`
}

func (l Logger) Zap() *zap.Logger {
	console := l.Console
	var filename string
	if l.Lumber != nil {
		filename = l.Lumber.Filename
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
	if console && filename != "" {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(l.Lumber), zapcore.AddSync(os.Stdout))
	} else if filename != "" {
		syncer = zapcore.AddSync(l.Lumber)
	} else {
		syncer = zapcore.AddSync(os.Stdout)
	}

	encoder := zapcore.NewConsoleEncoder(prod)

	level := zapcore.WarnLevel
	_ = level.Set(l.Level) // 就算设置失败还是默认值 WarnLevel
	core := zapcore.NewCore(encoder, syncer, level)

	return zap.New(core, zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel))
}
