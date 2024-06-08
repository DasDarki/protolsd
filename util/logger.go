package util

import "log"

type Logger struct {
	quiet   bool
	verbose bool
	debug   bool
}

func NewLogger(quiet bool, verbose bool, debug bool) *Logger {
	return &Logger{
		quiet:   quiet,
		verbose: verbose,
		debug:   debug,
	}
}

func (l *Logger) IsQuiet() bool {
	return l.quiet
}

func (l *Logger) IsVerbose() bool {
	return l.verbose
}

func (l *Logger) IsDebug() bool {
	return l.debug
}

func (l *Logger) Log(format string, args ...interface{}) {
	if !l.quiet {
		if l.verbose {
			log.Printf(format, args...)
		} else if l.debug {
			log.Printf(format, args...)
		}
	}
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.Log("WARNING: "+format, args...)
}

func (l *Logger) Err(format string, args ...interface{}) {
	l.Log("ERROR: "+format, args...)
}

func (l *Logger) Crit(format string, args ...interface{}) {
	l.Log("CRITICAL: "+format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.verbose {
		l.Log("INFO: "+format, args...)
	}
}
