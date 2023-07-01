package main

import (
	"WebGetterCLI/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting WebGetter CLI")
	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
