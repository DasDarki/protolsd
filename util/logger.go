package util

import "log"

type Logger interface {
	IsQuiet() bool
	IsVerbose() bool
	IsDebug() bool
	Log(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Err(format string, args ...interface{})
	Crit(format string, args ...interface{})
	Info(format string, args ...interface{})
}

type ConsoleLogger struct {
	quiet   bool
	verbose bool
	debug   bool
}

func NewLogger(quiet bool, verbose bool, debug bool) Logger {
	return &ConsoleLogger{
		quiet:   quiet,
		verbose: verbose,
		debug:   debug,
	}
}

func (l *ConsoleLogger) IsQuiet() bool {
	return l.quiet
}

func (l *ConsoleLogger) IsVerbose() bool {
	return l.verbose
}

func (l *ConsoleLogger) IsDebug() bool {
	return l.debug
}

func (l *ConsoleLogger) Log(format string, args ...interface{}) {
	if !l.quiet {
		if l.verbose {
			log.Printf(format, args...)
		} else if l.debug {
			log.Printf(format, args...)
		}
	}
}

func (l *ConsoleLogger) Warn(format string, args ...interface{}) {
	l.Log("WARNING: "+format, args...)
}

func (l *ConsoleLogger) Err(format string, args ...interface{}) {
	l.Log("ERROR: "+format, args...)
}

func (l *ConsoleLogger) Crit(format string, args ...interface{}) {
	l.Log("CRITICAL: "+format, args...)
}

func (l *ConsoleLogger) Info(format string, args ...interface{}) {
	if l.verbose {
		l.Log("INFO: "+format, args...)
	}
}
