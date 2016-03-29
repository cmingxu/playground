package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func init() {
}

func main() {
	var language string
	app := cli.NewApp()
	app.Name = "cli demo"
	app.Usage = "cli [OPTIONS] action"

	app.Action = func(ctx *cli.Context) {
		fmt.Println("action to take")
		fmt.Println(ctx.String("lang"))
		fmt.Println(language)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}
	app.Run(os.Args)
}
