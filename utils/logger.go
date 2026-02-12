package utils

import "fmt"

type Logger struct {
	Prefix string
}

func NewLogger() *Logger {
	return &Logger{
		Prefix: "[lazydocs]",
	}
}

func (l *Logger) Info(msg string, a ...any) {
	fmt.Printf("%s %s\n", l.Prefix, fmt.Sprintf(msg, a...))
}
