package main

import (
	"os"

	"github.com/kieranajp/victim/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
		Action: server.Start,
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("App crashed")
	}
}
