package handler

import (
	"net/http"

	"github.com/kieranajp/victim/pkg/driver"
	"github.com/kieranajp/victim/pkg/handler"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func StartWebhookMode(c *cli.Context) error {
	api := driver.NewSlackClient(c.String("slack-app-token"), c.String("slack-bot-token"))

	http.HandleFunc("/slack/events", handler.HandleWebhookVerification)

	http.HandleFunc("/slack/commands", handler.HandleSlashCommand)

	ih := &handler.InteractionHandler{API: api}
	http.HandleFunc("/slack/interactions", ih.HandleInteraction)

	log.Info().Msg("Server listening")
	http.ListenAndServe(":3000", nil)

	return nil
}
