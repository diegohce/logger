package logger

import (
	"log"
	"os"
	"strings"
)

// Logger logging facility object.
type Logger struct {
	stdout *log.Logger
	stderr *log.Logger
	debug  *log.Logger
}

// New creates a new Logger.
// prefix will be printed at the beginning of
// every line.
func New(prefix string) *Logger {

	var debugMode bool

	switch strings.ToLower(os.Getenv("DEBUG")) {
	case "true", "yes", "1":
		debugMode = true
	default:
		debugMode = false
	}

	l := &Logger{
		stdout: log.New(os.Stdout, prefix, log.LstdFlags|log.LUTC|log.Lmicroseconds),
		stderr: log.New(os.Stderr, prefix, log.LstdFlags|log.LUTC|log.Lmicroseconds),
		debug:  log.New(os.Stdout, prefix, log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile),
	}

	if !debugMode {
		devnull, _ := os.Open(os.DevNull)
		l.debug.SetOutput(devnull)
	}
	return l
}

// Info writes to stdout
func (l *Logger) Info() *log.Logger {
	return l.stdout
}

// Error writes to stderr
func (l *Logger) Error() *log.Logger {
	return l.stderr
}

// Debug writes to stdout if env variable DEBUG=1
// Debug writes to /dev/null.
func (l *Logger) Debug() *log.Logger {
	return l.debug
}
