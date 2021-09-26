package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const slackOAuthUrl = "https://slack.com/oauth/v2/authorize?client_id=%s&scope=%s&redirect_uri=%s"
const slackRedirectUrl = "https://f269-91-64-172-149.ngrok.io/slack/oauth/authorize"
const slackAccessUrl = "https://slack.com/api/oauth.v2.access"

var slackClient = &http.Client{Timeout: 10 * time.Second}
var requiredScopes = []string{
	"channels:read",
	"chat:write",
	"commands",
	"groups:read",
	"usergroups:read",
	"users:read",
}

func (o *OAuthHandler) Redirect(rw http.ResponseWriter, r *http.Request) {
	slackUri := fmt.Sprintf(
		slackOAuthUrl,
		o.ClientID,
		strings.Join(requiredScopes, ","),
		slackRedirectUrl,
	)

	http.Redirect(rw, r, slackUri, http.StatusFound)
}

func (o *OAuthHandler) Authorize(rw http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	resp, err := slackClient.PostForm(
		slackAccessUrl,
		url.Values{
			"code":          {code},
			"client_id":     {o.ClientID},
			"client_secret": {o.ClientSecret},
		},
	)

	if err != nil {
		log.Err(err).Msg("Failed to get access token")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Err(err).Msg("Invalid format for access token")
	}

	bodyStr := string(body)
	log.Info().Str("body", bodyStr).Msg("Access token response")

	teamID := gjson.Get(bodyStr, "team.id")
	teamName := gjson.Get(bodyStr, "team.name")
	token := gjson.Get(bodyStr, "access_token")

	log.Info().
		Str("team_id", teamID.String()).
		Str("team_name", teamName.String()).
		Str("token", token.String()).
		Msg("Access token retrieved")
}
