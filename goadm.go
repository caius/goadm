package goadm

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

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

// @private
func (g Client) exec(command string) ([]byte, error) {
	cmd := exec.Command(
		"ssh",
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "StrictHostKeyChecking=no",
		"-p", fmt.Sprintf("%d", g.Port),
		fmt.Sprintf("%s@%s", g.User, g.Host),
		command,
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	var outerr bytes.Buffer
	cmd.Stderr = &outerr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("err out: %q\n", outerr.String())
		log.Fatal(err)
	}

	return out.Bytes(), nil
}
