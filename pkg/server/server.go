package server

import (
	"net/http"

	"github.com/kieranajp/victim/pkg/driver"
	"github.com/kieranajp/victim/pkg/handler"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func Start(c *cli.Context) error {
	sh := &handler.SlackHandler{
		API: driver.NewSlackClient(
			c.String("slack-app-token"),
			c.String("slack-bot-token"),
		),
	}

	http.HandleFunc("/slack/events", handler.HandleWebhookVerification)
	http.HandleFunc("/slack/commands", sh.HandleSlashCommand)
	http.HandleFunc("/slack/interactions", sh.HandleInteraction)

	log.Info().Msg("Server listening")
	http.ListenAndServe(":3000", nil)

	return nil
}
