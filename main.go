package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harryzcy/ascheck/internal/macapp"
	"github.com/harryzcy/ascheck/internal/output"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "A cli app that check app's Apple Silicon support",
		Action: func(c *cli.Context) error {
			apps, err := macapp.GetAllApplications(nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			output.Table(apps)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
