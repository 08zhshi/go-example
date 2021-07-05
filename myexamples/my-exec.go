package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	args := []string{"-22", "-34"}
	env := os.Environ()

	err = syscall.Exec(cmd, args, env)
	fmt.Println("err=", err) // 这里不会输出东西，子进程占领父进程资源。
}
