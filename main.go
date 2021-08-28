package main

import (
	"log"
	"os"

	"github.com/kieranajp/victim/pkg/handler"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Victim",
		Usage: "Pick a name out of a hat.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "slack-app-token",
				Usage:    "Slack app token (for websockets)",
				EnvVars:  []string{"SLACK_APP_TOKEN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "slack-bot-token",
				Usage:    "Slack bot token (for @mentions)",
				EnvVars:  []string{"SLACK_BOT_TOKEN"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "socket",
				Usage:  "Start in socket mode",
				Action: handler.StartSocketMode,
			},
		},
		Action: handler.StartSocketMode,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
