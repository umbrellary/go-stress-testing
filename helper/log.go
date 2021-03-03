package helper

import (
	"os"
	"time"

	"github.com/fatih/color"
)

// Clog Clog
var Clog *LogAttributes

// LogAttributes Log Attributes
type LogAttributes struct {
	ServiceName string
}

func init() {
	Clog = &LogAttributes{
		ServiceName: "MAXIMO",
	}
}

// Infoln Info with ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Infoln(a ...interface{}) {
	color.New(color.FgHiRed).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"] ")
	color.New(color.FgGreen).Fprintln(os.Stdout, a...)
}

// Info Info without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Info(a ...interface{}) {
	color.New(color.FgGreen).Fprint(os.Stdout, a...)
}

// Infof Info without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Infof(format string, a ...interface{}) {
	color.New(color.FgGreen).Fprintf(os.Stdout, format, a...)
}

// Errorln Error with ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Errorln(a ...interface{}) {
	color.New(color.FgHiRed).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"] ")
	color.New(color.FgHiRed).Fprintln(os.Stdout, a...)
}

// Error Error without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Error(a ...interface{}) {
	color.New(color.FgHiRed).Fprint(os.Stdout, a...)
}

// Errorf Error without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Errorf(format string, a ...interface{}) {
	color.New(color.FgHiRed).Fprintf(os.Stdout, format, a...)
}

// Warnln Warn with ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Warnln(a ...interface{}) {
	color.New(color.FgHiRed).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"] ")
	color.New(color.FgYellow).Fprintln(os.Stdout, a...)
}

// Warn Warn without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Warn(a ...interface{}) {
	color.New(color.FgYellow).Fprint(os.Stdout, a...)
}

// Warnf Warn without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Warnf(format string, a ...interface{}) {
	color.New(color.FgYellow).Fprintf(os.Stdout, format, a...)
}

// Traceln Trace with ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Traceln(a ...interface{}) {
	color.New(color.FgHiRed).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"] ")
	color.New(color.FgCyan).Fprintln(os.Stdout, a...)
}

// Trace Trace without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Trace(a ...interface{}) {
	color.New(color.FgCyan).Fprint(os.Stdout, a...)
}

// Tracef Trace without ln.
// stdout bool 是否输出到控制台
func (la *LogAttributes) Tracef(format string, a ...interface{}) {
	color.New(color.FgCyan).Fprintf(os.Stdout, format, a...)
}
