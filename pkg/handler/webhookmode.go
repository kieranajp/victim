package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kieranajp/victim/pkg/driver"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

type InteractionPayload struct {
	Channel struct {
		ID string `json:"id"`
	} `json:"channel"`
	Actions []struct {
		Value string `json:"value"`
	} `json:"actions"`
}

func (p *InteractionPayload) ChannelID() string {
	return p.Channel.ID
}

func (p *InteractionPayload) GetUsers() (users []string) {
	for _, action := range p.Actions {
		users = append(users, strings.Split(action.Value, ",")...)
	}
	return
}

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
			Msg("Received Slack slash command")

		users := ExtractUsers(r.FormValue("text"))

		response := map[string]interface{}{
			"blocks": []slack.Block{
				slack.NewSectionBlock(
					&slack.TextBlockObject{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("Okay, I'll pick a victim from these users: %s", strings.Join(users, ", ")),
					},
					nil,
					slack.NewAccessory(
						slack.NewButtonBlockElement(
							"roll",
							strings.Join(users, ","),
							&slack.TextBlockObject{
								Type: slack.PlainTextType,
								Text: "Roll",
							},
						),
					),
				),
			}}

		responseJson, _ := json.Marshal(response)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(responseJson)
	})

	http.HandleFunc("/slack/interactions", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var p InteractionPayload
		json.Unmarshal([]byte(r.FormValue("payload")), &p)
		users := p.GetUsers()

		log.Info().
			Str("users", fmt.Sprintf("%+v\n", users)).
			Msg("Received Slack interaction")

		user := PickRandomUser(users)

		_, _, err := api.PostMessage(p.ChannelID(), slack.MsgOptionText(fmt.Sprintf("I have chosen: %s", user), false))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed posting message via Slack API")
		}
	})

	log.Info().Msg("Server listening")
	http.ListenAndServe(":3000", nil)

	return nil
}
