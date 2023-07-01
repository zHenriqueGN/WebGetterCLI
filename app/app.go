package app

import "github.com/urfave/cli"

// Returns the CLI app ready to execution
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebGetterCLI"
	app.Usage = "Search for IP addresses and servers"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for IP addresses on web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "gooogle.com.br",
				},
			},
			Action: searchIP,
		},
	}

	return app
}

func searchIP(c *cli.Context) {

}
