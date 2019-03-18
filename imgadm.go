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
		images = append(images, imagejsonToImage(json_image.ImageJSON))
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

	img := imagejsonToImage(parsedImage)
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
