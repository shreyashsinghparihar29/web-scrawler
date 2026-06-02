package main

import (
	"log"
	"os"

	"github.com/shreyashsinghparihar29/web-scrawler/cmd/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "web-scrawler",
		Description: "A lightweight web scraping and HTML data extraction tool powered by CSS selectors and JavaScript expressions",
		Commands:    []*cli.Command{},
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "extract",
		Description: "extract the required information from the specified URL using custom extractors",
		Action:      commands.Extractor(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "the target URL to extract data from",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "user-agent",
				Usage:   "the User-Agent header value",
				Aliases: []string{"ua"},
				Value:   "web-scrawler/fetch",
			},
			&cli.StringSliceFlag{
				Name:     "extract",
				Aliases:  []string{"x"},
				Required: true,
				Usage:    "extractor(s) to execute against the target URL in the form: -x key=script -x key2=script2",
			},
			&cli.BoolFlag{
				Name:  "return-body",
				Usage: "include the response body in the output",
				Value: false,
			},
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "serve",
		Description: "start the HTTP API server",
		Action:      commands.HTTPServer(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"a"},
				Usage:   "HTTP server listen address",
				Value:   ":8010",
			},
			&cli.BoolFlag{
				Name:    "logging",
				Aliases: []string{"l"},
				Value:   true,
				Usage:   "enable or disable access logging",
			},
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "shell",
		Description: "interactive JavaScript shell for debugging and testing extractors",
		Action:      commands.Shell(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Required: true,
				Usage:    "the page URL to inspect",
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
