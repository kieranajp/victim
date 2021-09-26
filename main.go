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
				Name:     "slack-client-id",
				Usage:    "Slack OAuth Client ID",
				EnvVars:  []string{"SLACK_CLIENT_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "slack-client-secret",
				Usage:    "Slack OAuth Client Secret",
				EnvVars:  []string{"SLACK_CLIENT_SECRET"},
				Required: true,
			},
			&cli.StringFlag{
				Name:        "listen-address",
				Usage:       "Host and port to listen on",
				EnvVars:     []string{"LISTEN_ADDRESS"},
				Value:       "127.0.0.1:3000",
				DefaultText: "127.0.0.1:3000",
			},
		},
		Action: server.Start,
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Exit")
	}
}
