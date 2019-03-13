package goadm

import (
	"encoding/json"
)

type Vmadm struct {
	Client
}

type ZonesJSON []struct {
	Manifest struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	} `json:"manifest"`
}

type Zone struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func (i Vmadm) ListZones() ([]Zone, error) {
	result, err := i.exec("vmadm lookup --json")
	if err != nil {
		return nil, err
	}

	var parsedZones ZonesJSON
	err = json.Unmarshal([]byte(result), &parsedZones)
	if err != nil {
		return nil, err
	}

	var zones []Zone
	for _, parsedZone := range parsedZones {
		zones = append(zones, Zone{
			Uuid: parsedZone.Manifest.Uuid,
			Name: parsedZone.Manifest.Name,
		})
	}

	return zones, nil
}
