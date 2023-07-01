package app

import "github.com/urfave/cli"

// Returns the CLI app ready to execution
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebGetterCLI"
	app.Usage = "Search for IP addresses and servers"
	return app
}
