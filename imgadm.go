package goadm

import (
	"encoding/json"
	"time"
)

type Imgadm struct {
	goadm Client
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
