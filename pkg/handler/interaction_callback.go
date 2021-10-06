package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/kieranajp/victim/pkg/database"
	"github.com/kieranajp/victim/pkg/driver"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
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

func HandleInteraction(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var p InteractionPayload
	json.Unmarshal([]byte(r.FormValue("payload")), &p)
	users := p.GetUsers()

	log.Info().
		Str("users", fmt.Sprintf("%+v\n", users)).
		Msg("Received Slack interaction")

	if len(users) == 0 {
		log.Error().Msg("No users in request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	accessToken, botToken, err := database.GetToken(database.New(), r.FormValue("team_id"))
	if err != nil || len(accessToken) < 1 {
		log.Fatal().
			Err(err).
			Str("team_id", r.FormValue("team_id")).
			Str("team_name", r.FormValue("team_domain")).
			Msg("Unknown team")
	}

	api := driver.NewSlackClient(accessToken, botToken)

	user := PickRandomUser(users)

	_, _, err = api.PostMessage(p.ChannelID(), slack.MsgOptionText(fmt.Sprintf("I have chosen: %s", user), false))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed posting message via Slack API")
	}
}

// PickRandomUser takes in a map of users and chooses a random one.
func PickRandomUser(users []string) string {
	log.Info().
		Str("Users", strings.Join(users, ",")).
		Msg("Picking random user")

	rand.Seed(time.Now().Unix())
	user := users[rand.Intn(len(users))]
	return user
}
