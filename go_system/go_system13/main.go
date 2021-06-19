package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
)

func main() {
	// fmt.Printf("プロセスID: %d\n", os.Getpid())
	// fmt.Printf("プロセスID: %d\n", os.Getppid())

	// sid, _ := syscall.Getsid(os.Getpid())
	// fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", syscall.Getpgrp(), sid)

	// fmt.Printf("ユーザーID: %d\n", os.Getuid())
	// fmt.Printf("グループID: %d\n", os.Getgid())
	// groups, _ := os.Getgroups()
	// fmt.Printf("サブグループID: %v\n", groups)
	// fmt.Println(os.Getwd())

	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	fmt.Printf("parent pid: %d name: '%s' cmd: '%s'\n", p.Pid, name, cmd)
}
