package slog

import (
	"log"
	"os"
	"sync"
)

var debugFlag = func() bool {
	flag := os.Getenv("DEBUG")
	return flag != "" && flag != "0" && flag != "false"
}()

type (
	Logger interface {
		// Println calls l.Output to print to the logger.
		// Arguments are handled in the manner of `fmt.Println`.
		Println(v ...interface{})
		// Printf calls l.Output to print to the logger.
		// Arguments are handled in the manner of `fmt.Printf`.
		Printf(format string, v ...interface{})
		// SetPrefix sets the output prefix for the logger.
		SetPrefix(v string)
	}

	logger struct {
		src *log.Logger
		mu  sync.Mutex
	}
)

// NewLogger returns a new implementation of `Logger`.
func NewLogger(prefix string) Logger {
	return &logger{
		src: log.New(os.Stdout, prefix, 0),
		mu:  sync.Mutex{},
	}
}

func (l *logger) Println(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if debugFlag {
		l.src.Println(v...)
	}
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if debugFlag {
		l.src.Printf(format, v...)
	}
}

func (l *logger) SetPrefix(v string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.src.SetPrefix(v)
}
