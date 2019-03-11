package main

import (
	"log"
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

type Imgadm struct {
	goadm Goadm
}

type Image struct {
	uuid      string
	name      string
	installed bool
}

func (i Imgadm) ListImages() ([Image], error) {
	result, err := goadm.Exec("imgadm list -j")
  if err != nil {
    return nil, err
  }

  return JSON.parse(result), nil
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

	for image := range imgadm.ListImages() {
		log.Info(image.Name)
	}

}
