/*
@Time : 2022/5/21 11:02
@Author : lianyz
@Description :
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

const mycontaierMemoryCgroups = "/sys/fs/cgroup/memory/mycontainer"

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Println("[main]", "pid:", os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Println("[exe]", "pid:", os.Getpid())

	setCGroup()
	must(syscall.Sethostname([]byte("mycontainer")))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func setCGroup() {
	os.Mkdir(mycontaierMemoryCgroups, 0755)
	writeCGroup("memory.limit_in_bytes", "100M")
	writeCGroup("notify_on_release", "1")
	writeCGroup("tasks", strconv.Itoa(os.Getpid()))
}

func writeCGroup(fileName string, data string) {
	must(ioutil.WriteFile(filepath.Join(mycontaierMemoryCgroups,
		fileName), []byte(data), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
