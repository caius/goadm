package goadm

import (
	"encoding/json"
	"fmt"
)

type Vmadm struct {
	Client
}

type ZonesJSON []struct {
	ZoneJSON
}

type ZoneJSON struct {
	Autoboot          bool   `json:"autoboot"`
	Brand             string `json:"brand"`
	CpuShares         int    `json:"cpu_shares"`
	DnsDomain         string `json:"dns_domain"`
	ImageUuid         string `json:"image_uuid"`
	MaxLockedMemory   int    `json:"max_locked_memory"`
	MaxLwps           int    `json:"max_lwps"`
	MaxPhysicalMemory int    `json:"max_physical_memory"`
	MaxSwap           int    `json:"max_swap"`
	Name              string `json:"name"`
	State             string `json:"zone_state"`
	Uuid              string `json:"uuid"`
	ZfsIoPriority     int    `json:"zfs_io_priority"`
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
			Uuid: parsedZone.Uuid,
			Name: parsedZone.Name,
		})
	}

	return zones, nil
}

func (i Vmadm) GetZone(uuid string) (*Zone, error) {
	result, err := i.exec(fmt.Sprintf("vmadm get %s", uuid))
	if err != nil {
		return nil, err
	}

	var parsedZone ZoneJSON
	err = json.Unmarshal([]byte(result), &parsedZone)

	return &Zone{
		Uuid: parsedZone.Uuid,
		Name: parsedZone.Name,
	}, nil
}
