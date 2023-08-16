package lib

import (
	"log"
	"os"
)

type Logger struct {
	info    *log.Logger
	warn    *log.Logger
	err     *log.Logger
	dbg     *log.Logger
	verbose bool
}

func (l Logger) Info(format string, args ...any) {
	l.info.Printf(format, args...)
}

func (l Logger) Warn(format string, args ...any) {
	l.warn.Printf(format, args...)
}

func (l Logger) Err(format string, args ...any) {
	l.err.Printf(format, args...)
}

func (l Logger) ErrWrap(err error) {
	if err != nil {
		l.err.Printf("%v", err)
	}
}

func (l Logger) Dbg(format string, args ...any) {
	if l.verbose {
		l.dbg.Printf(format, args...)
	}
}

var Log Logger = Logger{
	info: log.New(os.Stdin, "", log.Ldate|log.Ltime|log.Lmsgprefix),
	warn: log.New(os.Stderr, "Warning: ", log.Ldate|log.Ltime|log.Lmsgprefix),
	err:  log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lmsgprefix),
	dbg:  log.New(os.Stderr, "(Debug) ", log.Ldate|log.Ltime|log.Lmsgprefix),
}

func SetLogDebug(flag bool) {
	Log.verbose = flag
}

func SetLogFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	Log.info.SetOutput(file)
	Log.warn.SetOutput(file)
	Log.err.SetOutput(file)
	Log.dbg.SetOutput(file)
	return nil
}
