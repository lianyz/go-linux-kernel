/*
@Time : 2022/5/21 11:02
@Author : lianyz
@Description :
*/

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	must(cmd.Run(syscall.Sethostname([]byte("container1"))))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
