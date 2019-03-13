package goadm

import (
	"encoding/json"
	"time"
)

type Imgadm struct {
	Client
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
