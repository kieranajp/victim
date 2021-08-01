package driver

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func NewSlackClient(appToken, botToken string) (*slack.Client, *socketmode.Client) {
	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(appToken),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	return api, client
}
