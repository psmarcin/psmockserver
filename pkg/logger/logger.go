package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/kataras/golog"
)

func Setup() {
	golog.SetTimeFormat("")
	golog.Handle(func(l *golog.Log) bool {
		prefix := golog.GetTextForLevel(l.Level, true)

		_, line := getCaller()
		message := fmt.Sprintf("%s %s [%s:%d] %s",
			prefix, l.FormatTime(), "", line, l.Message)

		if l.NewLine {
			message += "\n"
		}

		fmt.Print(message)
		return true
	})
}

// https://golang.org/doc/go1.9#callersframes
func getCaller() (string, int) {
	var pcs [10]uintptr
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()

		if !strings.HasSuffix(frame.File, "github.com/kataras/golog") && frame.Func.Name() != "main.getCaller" {
			return frame.File, frame.Line
		}

		if !more {
			break
		}
	}

	return "?", 0
}
