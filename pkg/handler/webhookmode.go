package handler

import (
	"fmt"
	"net/http"

	"github.com/kieranajp/victim/pkg/driver"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func StartWebhookMode(c *cli.Context) error {
	api, _ := driver.NewSlackClient(c.String("slack-app-token"), c.String("slack-bot-token"))

	http.HandleFunc("/slack/events", driver.WebhookVerifier)

	http.HandleFunc("/slack/commands", func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid incoming webhook")
		}

		log.Info().
			Str("payload", fmt.Sprintf("%+v\n", r.PostForm)).
			Msg("Received Slack webhook")

		users := ExtractUsers(r.FormValue("text"))
		user := PickRandomUser(users)

		_, _, err = api.PostMessage(r.FormValue("channel_id"), slack.MsgOptionText(fmt.Sprintf("I have chosen: <%s>", user), false))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed posting message via Slack API")
		}
	})

	log.Info().Msg("Server listening")
	http.ListenAndServe(":3000", nil)

	return nil
}
