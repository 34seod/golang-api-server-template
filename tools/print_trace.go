package tools

import (
	"log"
	"runtime"
)

func PrintTrace() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	for _, p := range pc {
		f := runtime.FuncForPC(p)
		file, line := f.FileLine(p)
		log.Printf("\033[31m%s:%d\n", file, line)
	}
}
