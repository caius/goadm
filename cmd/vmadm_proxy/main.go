package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caius/goadm"
	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Name = "vmadm_proxy"
	app.Usage = "Proxy for vmadm on another host"

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List zones",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "SmartOS Host to run against",
				},
				cli.StringFlag{
					Name:  "user",
					Usage: "User to login as",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "SSH Port on host",
				},
			},
			Action: func(c *cli.Context) error {
				client := goadm.NewClient(c.String("host"), c.String("user"), c.Int("port"))

				vmadm := client.Vmadm()
				zones, err := vmadm.ListZones()
				if err != nil {
					return err
				}

				fmt.Printf("%s\t%s\n", "UUID", "Name")

				for _, zone := range zones {
					fmt.Printf("%s\t%s\n", zone.Uuid, zone.Name)
				}

				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get information on a single zone",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "SmartOS Host to run against",
				},
				cli.StringFlag{
					Name:  "user",
					Usage: "User to login as",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "SSH Port on host",
				},
			},
			Action: func(c *cli.Context) error {
				client := goadm.NewClient(c.String("host"), c.String("user"), c.Int("port"))

				vmadm := client.Vmadm()
				zone, err := vmadm.GetZone(c.Args().Get(0))
				if err != nil {
					return err
				}

				fmt.Printf("%s\t%s\n", "UUID", "Name")
				fmt.Printf("%s\t%s\n", zone.Uuid, zone.Type)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
