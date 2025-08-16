package util

import "fmt"

type Logger struct {
	Tag string
}

func (l *Logger) Log(format string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", l.Tag, fmt.Sprintf(format, args...))
}
