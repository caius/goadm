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
	app.Name = "imgadm_proxy"
	app.Usage = "Proxy for imgadm on another host"

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List installed images",
			Action: func(c *cli.Context) error {
				client := goadm.NewClient("127.0.0.1", "root", 2022)

				imgadm := client.Imgadm()
				images, err := imgadm.ListImages()
				if err != nil {
					return err
				}

				fmt.Printf("%s\t%s\t%s\n", "UUID", "Name", "Version")

				for _, image := range images {
					fmt.Printf("%s\t%s\t%s\n", image.Uuid, image.Name, image.Version)
				}

				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get information on a single image",
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

				imgadm := client.Imgadm()
				image, err := imgadm.GetImage(c.Args().Get(0))
				if err != nil {
					return err
				}

				fmt.Printf("%s\n", image.Uuid)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
