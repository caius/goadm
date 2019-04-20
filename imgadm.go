package goadm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

// Relevant documentation from imgadm
//
// Errors: https://github.com/joyent/smartos-live/blob/master/src/img/lib/errors.js

type Imgadm struct {
	Client
}

type ImagesJSON []struct {
	ImageJSON
}

type ImageJSON struct {
	Manifest struct {
		Uuid        string    `json:"uuid"`
		Name        string    `json:"name"`
		Os          string    `json:"os"`
		PublishedAt time.Time `json:"published_at"`
		Type        string    `json:"type"`
		Urn         string    `json:"urn"`
		Version     string    `json:"version"`
	} `json:"manifest"`
}

type Image struct {
	Uuid        string
	Name        string
	Os          string
	PublishedAt time.Time
	Type        string
	Urn         string
	Version     string
}

func (i Imgadm) ListImages() ([]Image, error) {
	result := i.exec("imgadm list --json")

	log.Printf("GOADM ListImages result=%+v\n", result)

	if result.ExitCode != 0 {
		// Uh oh, our command failed
		log.Printf("GOADM ListImages exec failed: %+v\n", result)
		return nil, errors.New("Error listing images")
	}

	var parsedImages ImagesJSON
	err := json.Unmarshal([]byte(result.Stdout), &parsedImages)
	if err != nil {
		// Uh oh, our command failed
		log.Printf("GOADM ListImages unmarshal failed: %+v\n", err)
		return nil, errors.New("Error listing images")
	}

	var images []Image
	for _, json_image := range parsedImages {
		images = append(images, imagejsonToImage(json_image.ImageJSON))
	}

	return images, nil
}

// Gets installed image
func (i Imgadm) GetImage(uuid string) (*Image, error) {
	log.Printf("GOADM GetImage uuid=%s\n", uuid)
	result := i.exec(fmt.Sprintf("imgadm get %s", uuid))

	log.Printf("GOADM result=%+v\n", result)

	if result.ExitCode != 0 {
		log.Printf("GOADM GetImage exec failed: %+v\n", result)
		// Uh oh, our command failed, assume not found for now
		return nil, errors.New("Image not found")
	}

	var parsedImage ImageJSON
	err := json.Unmarshal([]byte(result.Stdout), &parsedImage)
	if err != nil {
		log.Printf("GOADM GetImage unmarshal failed: %+v\n", err)
		// Uh oh, our command failed, assume not found for now
		return nil, errors.New("Image not found")
	}

	img := imagejsonToImage(parsedImage)
	log.Printf("GOADM img %+v\n", img)
	return &img, nil
}

func imagejsonToImage(data ImageJSON) Image {
	return Image{
		Uuid:        data.Manifest.Uuid,
		Name:        data.Manifest.Name,
		Os:          data.Manifest.Os,
		PublishedAt: data.Manifest.PublishedAt,
		Type:        data.Manifest.Type,
		Urn:         data.Manifest.Urn,
		Version:     data.Manifest.Version,
	}
}
