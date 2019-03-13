package main

import (
	"github.com/caius/goadm"
	"log"
)

func main() {
	client := goadm.NewClient("127.0.0.1", "root", 2022)

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
