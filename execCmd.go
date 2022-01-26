package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func showTime() {
	fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
}

func setupCmd(cmdName string, cmdArgs ...string) *exec.Cmd {
	cmd := exec.Command(cmdName, cmdArgs...)

	ro, wo := io.Pipe()
	cmd.Stdout = wo

	re, we := io.Pipe()
	cmd.Stderr = we

	go io.Copy(os.Stdout, ro)
	go io.Copy(os.Stderr, re)

	return cmd
}

func main() {
	showTime()
	defer showTime()

	fmt.Println("starting cmd")
	fmt.Println(setupCmd("cmd.exe", "/C", os.Args[1]).Run())
	fmt.Println("ran")
}
