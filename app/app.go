package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the CLI app ready to execution
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebGetterCLI"
	app.Usage = "Search for IP addresses and servers on web"
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IP addresses on web",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Search for servers on web",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

// Search for IPs to a given host
func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Search for servers to a given host
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
