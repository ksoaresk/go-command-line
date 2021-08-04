package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Retorna aplicação para ser executada.
func RunApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Application for command line"
	app.Usage = "Search IPs and NS"

	app.Commands = []cli.Command{
		{
			Name:    "ips",
			Aliases: []string{"i"},
			Usage:   "go run main.go ip|i --host {Your_Host.com[ .br ]}",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "hospedandoweb.com.br",
				},
			},
			Action: buscarIps,
		},
		{
			Name:    "servers",
			Aliases: []string{"s"},
			Usage:   "go run main.go servers|s --host {Your_Host.com[ .br ]}",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "hospedandoweb.com.br",
				},
			},
			Action: buscarHosts,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarHosts(c *cli.Context) {
	host := c.String("host")

	nameServers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ns := range nameServers {
		fmt.Println(ns.Host)
	}
}
