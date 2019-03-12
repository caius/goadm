package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Goadm struct {
	Host string
	User string
	Port int
}

func (g Goadm) Imgadm() (Imgadm, error) {
	imgadm := Imgadm{goadm: g}
	return imgadm, nil
}

func (g Goadm) Exec(command string) (string, error) {
	cmd := exec.Command("./exe/wrapper.sh", g.Host, g.User, fmt.Sprintf("%d", g.Port), "--", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var outerr bytes.Buffer
	cmd.Stderr = &outerr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("err out: %q\n", outerr.String())
		log.Fatal(err)
	}

	fmt.Printf("cmd out: %q\n", out.String())

	return "", nil
}

type Imgadm struct {
	goadm Goadm
}

type Image struct {
	Uuid      string
	Name      string
	Installed bool
}

func (i Imgadm) ListImages() ([]Image, error) {
	result, err := i.goadm.Exec("imgadm list -j")
	if err != nil {
		return nil, err
	}

	log.Printf("%s\n", result)

	// return JSON.parse(result), nil
	var images [1]Image
	images[0] = Image{
		Uuid: "97a044cd-10c3-4ba7-9fe4-933481e3474d",
		Name: "debian9",
	}
	return images[:], nil
}

func main() {
	client := Goadm{
		Host: "127.0.0.1",
		User: "root",
		Port: 2022,
	}

	imgadm, err := client.Imgadm()
	if err != nil {
		log.Fatal(err)
	}

	images, err := imgadm.ListImages()
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range images {
		log.Printf("%s\t%s\n", image.Uuid, image.Name)
	}
}
