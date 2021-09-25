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

	http.HandleFunc("/healthz", handler.WithLogging(handler.Healthz))
	http.HandleFunc("/slack/events", handler.WithLogging(handler.HandleWebhookVerification))
	http.HandleFunc("/slack/commands", handler.WithLogging(sh.HandleSlashCommand))
	http.HandleFunc("/slack/interactions", handler.WithLogging(sh.HandleInteraction))

	log.Info().
		Str("listen_address", c.String("listen-address")).
		Msg("Server listening")

	return http.ListenAndServe(c.String("listen-address"), nil)
}
