package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	w := io.Writer(os.Stdout)
	l := log.New(w, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(w, "\033[32mDEBUG:\033[0m", l.Flags()),
		info:   log.New(w, "\033[34mINFO:\033[0m ", l.Flags()),
		warn:   log.New(w, "\033[33mWARNING:\033[0m ", l.Flags()),
		err:    log.New(w, "\033[31mERROR:\033[0m ", l.Flags()),
		writer: w,
	}
}

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...any) {
	l.warn.Println(v...)
}
func (l *Logger) Err(v ...any) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(f string, v ...any) {
	l.debug.Printf(f, v...)
}
func (l *Logger) Infof(f string, v ...any) {
	l.info.Printf(f, v...)
}
func (l *Logger) Warnf(f string, v ...any) {
	l.warn.Printf(f, v...)
}
func (l *Logger) Errf(f string, v ...any) {
	l.err.Printf(f, v...)
}
