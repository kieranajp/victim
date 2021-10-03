package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
)

func (h *SlackHandler) HandleSlashCommand(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid incoming webhook")
	}

	log.Info().
		Str("payload", fmt.Sprintf("%+v\n", r.PostForm)).
		Msg("Received Slack slash command")

	users := ExtractUsers(r.FormValue("text"))
	users, err = ResolveUserGroups(users, h.API)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to resolve user groups")
	}

	if len(users) == 0 {
		log.Info().Msg("No users found")
		rw.Write([]byte("I couldn't find any matching users to victimise!"))
		return
	}

	responseJson := GenerateResponse(users)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(responseJson)
}

// ExtractUsers takes in the incoming text from a Slack command, and finds all the linkable @user entities within.
func ExtractUsers(text string) []string {
	r := regexp.MustCompile(`<(@[^<|>]*)[\|>]`)
	m := r.FindAllStringSubmatch(text, -1)

	var users []string
	for _, v := range m {
		users = append(users, fmt.Sprintf("<%s>", v[1]))
	}
	return users
}

func ExtractExclusions(text string) []string {
	r := regexp.MustCompile(`!<(@[^<|>]*)[\|>]`)
	m := r.FindAllStringSubmatch(text, -1)

	var exclusions []string
	for _, v := range m {
		exclusions = append(exclusions, strings.TrimPrefix(v[0], "!"))
	}
	return exclusions
}

func ResolveExclusions(users, exclusions []string) []string {
	resolved := make([]string, 0)
	for _, user := range users {
		included := true
		for _, exclusion := range exclusions {
			if exclusion == user {
				included = false
			}
		}
		if included {
			resolved = append(resolved, user)
		}
	}
	return resolved
}

func GenerateResponse(users []string) []byte {
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

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal response to JSON")
	}

	return responseJson
}
