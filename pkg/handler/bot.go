package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/kieranajp/victim/pkg/driver"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"github.com/urfave/cli/v2"
)

func StartBot(c *cli.Context) error {
	api, client := driver.NewSlackClient(c.String("slack-app-token"), c.String("slack-bot-token"))

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")
			case socketmode.EventTypeConnectionError:
				fmt.Println("Connection failed. Retrying later...")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}

				client.Debugf("Slash command received: %+v", cmd)

				users, _ := ResolveUserGroups(FindMentions(cmd.Text), api)

				payload := map[string]interface{}{
					"blocks": []slack.Block{
						slack.NewSectionBlock(
							&slack.TextBlockObject{
								Type: slack.MarkdownType,
								Text: fmt.Sprintf("Found these users: %s", strings.Join(users, ", ")),
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

				client.Ack(*evt.Request, payload)
			case socketmode.EventTypeInteractive:
				callback, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}

				fmt.Printf("Interaction received: %+v\n", callback)

				var payload interface{}

				switch callback.Type {
				case slack.InteractionTypeBlockActions:
					user := PickRandomUser(callback.ActionCallback.BlockActions[0].Value)
					_, _, err := api.PostMessage(callback.Channel.GroupConversation.ID, slack.MsgOptionText(fmt.Sprintf("I have chosen: %s", user), false))
					if err != nil {
						fmt.Printf("failed posting message: %v", err)
					}
				default:

				}

				client.Ack(*evt.Request, payload)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	client.Run()
	return nil
}
