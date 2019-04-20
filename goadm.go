package goadm

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type ExecResult struct {
	ExitCode int
	Stderr   string
	Stdout   string
}

type Client struct {
	Host string
	User string
	Port int
}

func NewClient(host string, user string, port int) Client {
	return Client{
		Host: host,
		User: user,
		Port: port,
	}
}

func (g Client) Imgadm() Imgadm {
	return Imgadm{Client: g}
}

func (g Client) Vmadm() Vmadm {
	return Vmadm{Client: g}
}

func (g Client) exec(command string) ExecResult {
	cmd := exec.Command(
		"ssh",
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "StrictHostKeyChecking=no",
		"-p", fmt.Sprintf("%d", g.Port),
		fmt.Sprintf("%s@%s", g.User, g.Host),
		command,
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Exec error: %s\n", err)
	}

	result := ExecResult{
		Stderr: stderr.String(),
		Stdout: stdout.String(),
	}

	return result
}
