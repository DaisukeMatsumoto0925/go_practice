package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, line, ok := runtime.Caller(0)

	if ok {
		name := filepath.Base(file)
		fmt.Printf("file: %s, line: %d\n", name, line)
	}
}
