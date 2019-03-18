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
	Type              string `json:"type"`
}

type Zone struct {
	Autoboot          bool
	Brand             string
	CpuShares         int
	DnsDomain         string
	ImageUuid         string
	MaxLockedMemory   int
	MaxLwps           int
	MaxPhysicalMemory int
	MaxSwap           int
	Name              string
	State             string
	Type              string
	Uuid              string
	ZfsIoPriority     int
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
		zones = append(zones, zonejsonToZone(parsedZone.ZoneJSON))
	}

	return zones, nil
}

func (i Vmadm) GetZone(uuid string) (*Zone, error) {
	// Lookup always returns an array, but we want to error if more than one VM matches
	// UUIDs *are* unique
	result, err := i.exec(fmt.Sprintf("vmadm lookup --json -1 uuid=%s", uuid))
	if err != nil {
		return nil, err
	}

	// Parse out the zones (there is only one)
	var parsedZones ZonesJSON
	err = json.Unmarshal([]byte(result), &parsedZones)
	if err != nil {
		return nil, err
	}

	// Return the one zone from the array
	zone := zonejsonToZone(parsedZones[0].ZoneJSON)
	return &zone, nil
}

func zonejsonToZone(data ZoneJSON) Zone {
	return Zone{
		Uuid:              data.Uuid,
		Autoboot:          data.Autoboot,
		Brand:             data.Brand,
		CpuShares:         data.CpuShares,
		DnsDomain:         data.DnsDomain,
		ImageUuid:         data.ImageUuid,
		MaxLockedMemory:   data.MaxLockedMemory,
		MaxLwps:           data.MaxLwps,
		MaxPhysicalMemory: data.MaxPhysicalMemory,
		MaxSwap:           data.MaxSwap,
		Name:              data.Name,
		State:             data.State,
		ZfsIoPriority:     data.ZfsIoPriority,
		Type:              data.Type,
	}
}
