package handler

import (
	"net/http"

	"github.com/slack-go/slack"
)

type SlackHandler struct {
	API *slack.Client
}

type OAuthHandler struct {
	ClientID     string
	ClientSecret string
}

func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
