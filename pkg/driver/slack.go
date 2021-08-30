package driver

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

func NewSlackClient(appToken, botToken string) *slack.Client {
	return slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(appToken),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)
}
