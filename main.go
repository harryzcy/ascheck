package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harryzcy/ascheck/internal/macapp"
	"github.com/harryzcy/ascheck/internal/output"
	"github.com/harryzcy/ascheck/internal/remotecheck"
	"github.com/urfave/cli/v2"
)

// handleErr prints error and calls os.Exit(1) if err is not nil.
func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	app := &cli.App{
		Usage:           "A cli app that check app's Apple Silicon support",
		Version:         "0.2.0",
		HideHelpCommand: true,
		UsageText:       "ascheck [global options]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "output in json format",
			},
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				EnvVars: []string{"LANG"},
				Usage:   "language for displaying app names",
			},
		},
		Action: func(c *cli.Context) error {
			err := remotecheck.Init()
			handleErr(err)

			apps, err := macapp.GetAllApplications(nil)
			handleErr(err)

			if c.Bool("json") {
				output.JSON(apps)
			} else {
				output.Table(apps)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
