package main

import (
	"fmt"
	"log"

	"github.com/caius/goadm"
)

func main() {
	client := goadm.NewClient("127.0.0.1", "root", 2022)

	vmadm, err := client.Vmadm()
	if err != nil {
		log.Fatal(err)
	}

	zones, err := vmadm.ListZones()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d zones\n", len(zones))

	for _, zone := range zones {
		fmt.Printf("%s\t%s\n", zone.Uuid, zone.Name)
	}
}
