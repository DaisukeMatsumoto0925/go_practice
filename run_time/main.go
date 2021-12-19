package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	err := func01()
	if err != nil {
		fmt.Println(err)
	}
	err = check("c", "d")
	if err != nil {
		fmt.Println(err)
	}
}

func check(a, b string) error {
	if a == b {
		return nil
	}

	msg := fmt.Sprintf(`check error. "%s" != "%s"`, a, b)
	return makeError(msg, 2)
}

func func01() error {
	return check("a", "b")
}

func makeError(msg string, skip int) error {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		name := filepath.Base(file)
		return fmt.Errorf(`%s; file: %s, line: %d, func: %s`,
			msg, name, line, runtime.FuncForPC(pc).Name())
	}
	return fmt.Errorf("%s", msg)
}
