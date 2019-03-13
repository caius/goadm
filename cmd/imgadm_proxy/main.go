package main

import (
	"fmt"
	"log"

	"github.com/caius/goadm"
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

	fmt.Printf("Found %d images\n", len(images))

	for _, image := range images {
		fmt.Printf("%s\t%s\n", image.Uuid, image.Name)
	}
}
