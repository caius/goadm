package goadm

import (
	"encoding/json"
	"fmt"
	"time"
)

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
		Version     string    `json:"version"`
		PublishedAt time.Time `json:"published_at"`
		Type        string    `json:"type"`
		Os          string    `json:"os"`
		Urn         string    `json:"urn"`
	} `json:"manifest"`
}

type Image struct {
	Uuid    string
	Name    string
	Version string
}

func (i Imgadm) ListImages() ([]Image, error) {
	result, err := i.exec("imgadm list --json")
	if err != nil {
		return nil, err
	}

	var parsedImages ImagesJSON
	err = json.Unmarshal([]byte(result), &parsedImages)
	if err != nil {
		return nil, err
	}

	var images []Image
	for _, json_image := range parsedImages {
		data := json_image.Manifest
		images = append(images, Image{
			Uuid:    data.Uuid,
			Name:    data.Name,
			Version: data.Version,
		})
	}

	return images, nil
}

// Gets installed image
func (i Imgadm) GetImage(uuid string) (*Image, error) {
	result, err := i.exec(fmt.Sprintf("imgadm get %s", uuid))
	if err != nil {
		return nil, err
	}

	var parsedImage ImageJSON
	err = json.Unmarshal(result, &parsedImage)
	if err != nil {
		return nil, err
	}

	data := parsedImage.Manifest
	return &Image{
		Uuid:    data.Uuid,
		Name:    data.Name,
		Version: data.Version,
	}, nil

}
