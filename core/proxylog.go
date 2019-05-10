package core

import (
	"fmt"

	logger "github.com/d2r2/go-logger"
	"github.com/davecgh/go-spew/spew"
)

type WriteLine func(line string) error

// ProxyLog is used to substitute regular log console output
// with output to the file, either to the GUI widget.
// ProxyLog implements logger.PackageLog interface which
// provide regular log methods.
type ProxyLog struct {
	parent      logger.PackageLog
	packageName string
	packageLen  int
	timeFormat  string

	customWriteLine WriteLine
	customLogLevel  logger.LogLevel
}

// Static cast to verify that type implement specific interface
var _ logger.PackageLog = &ProxyLog{}

func NewProxyLog(parent logger.PackageLog, packageName string, packageLen int,
	timeFormat string, writeLine WriteLine, customLogLevel logger.LogLevel) *ProxyLog {

	v := &ProxyLog{parent: parent, packageName: packageName, packageLen: packageLen,
		timeFormat: timeFormat, customLogLevel: customLogLevel,
		customWriteLine: writeLine}
	return v
}

func (v *ProxyLog) getFormat() logger.FormatOptions {
	options := logger.FormatOptions{TimeFormat: v.timeFormat,
		LevelLength: logger.LevelShort, PackageLength: v.packageLen}
	return options
}

func (v *ProxyLog) Printf(level logger.LogLevel, format string, args ...interface{}) {
	if v.parent != nil {
		v.parent.Printf(level, format, args...)
	}
	if v.customWriteLine != nil && level <= v.customLogLevel {
		msg := spew.Sprintf(format, args...)
		packageName := v.packageName
		out := logger.FormatMessage(v.getFormat(), level, packageName, msg, false)
		err := v.customWriteLine(out + fmt.Sprintln())
		if err != nil {
			v.parent.Fatal(err)
		}
	}
}

func (v *ProxyLog) Print(level logger.LogLevel, args ...interface{}) {
	if v.parent != nil {
		v.parent.Print(level, args...)
	}
	if v.customWriteLine != nil && level <= v.customLogLevel {
		msg := fmt.Sprint(args...)
		packageName := v.packageName
		out := logger.FormatMessage(v.getFormat(), level, packageName, msg, false)
		err := v.customWriteLine(out + fmt.Sprintln())
		if err != nil {
			v.parent.Fatal(err)
		}
	}
}

func (v *ProxyLog) Debugf(format string, args ...interface{}) {
	v.Printf(logger.DebugLevel, format, args...)
}

func (v *ProxyLog) Debug(args ...interface{}) {
	v.Print(logger.DebugLevel, args...)
}

func (v *ProxyLog) Infof(format string, args ...interface{}) {
	v.Printf(logger.InfoLevel, format, args...)
}

func (v *ProxyLog) Info(args ...interface{}) {
	v.Print(logger.InfoLevel, args...)
}

func (v *ProxyLog) Notifyf(format string, args ...interface{}) {
	v.Printf(logger.NotifyLevel, format, args...)
}

func (v *ProxyLog) Notify(args ...interface{}) {
	v.Print(logger.NotifyLevel, args...)
}

func (v *ProxyLog) Warningf(format string, args ...interface{}) {
	v.Printf(logger.WarnLevel, format, args...)
}

func (v *ProxyLog) Warnf(format string, args ...interface{}) {
	v.Printf(logger.WarnLevel, format, args...)
}

func (v *ProxyLog) Warning(args ...interface{}) {
	v.Print(logger.WarnLevel, args...)
}

func (v *ProxyLog) Warn(args ...interface{}) {
	v.Print(logger.WarnLevel, args...)
}

func (v *ProxyLog) Errorf(format string, args ...interface{}) {
	v.Printf(logger.ErrorLevel, format, args...)
}

func (v *ProxyLog) Error(args ...interface{}) {
	v.Print(logger.ErrorLevel, args...)
}

func (v *ProxyLog) Panicf(format string, args ...interface{}) {
	v.Printf(logger.PanicLevel, format, args...)
}

func (v *ProxyLog) Panic(args ...interface{}) {
	v.Print(logger.PanicLevel, args...)
}

func (v *ProxyLog) Fatalf(format string, args ...interface{}) {
	v.Printf(logger.FatalLevel, format, args...)
}

func (v *ProxyLog) Fatal(args ...interface{}) {
	v.Print(logger.FatalLevel, args...)
}
