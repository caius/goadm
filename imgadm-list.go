package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"
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

	return out.String(), nil
}

type Imgadm struct {
	goadm Goadm
}

type ImagesJSON []struct {
	Manifest struct {
		Uuid        string    `json:"uuid"`
		Name        string    `json:"name"`
		Version     string    `json:"version"`
		PublishedAt time.Time `json:"published_at"`
		Type        string    `json:"type"`
		Os          string    `json:"os"`
		Urn         string    `json:"urn"`
	} `json:"manifest"`
}

type Image struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func (i Imgadm) ListImages() ([]Image, error) {
	result, err := i.goadm.Exec("imgadm list -j")
	if err != nil {
		return nil, err
	}

	var parsedImages ImagesJSON
	err = json.Unmarshal([]byte(result), &parsedImages)
	if err != nil {
		return nil, err
	}

	var images []Image
	for _, image_j := range parsedImages {
		images = append(images, Image{
			Uuid: image_j.Manifest.Uuid,
			Name: image_j.Manifest.Name,
		})
	}

	return images, nil
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
