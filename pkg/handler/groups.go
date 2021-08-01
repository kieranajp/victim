package handler

import (
	"strings"

	"github.com/slack-go/slack"
)

func FindMentions(payload string) []string {
	input := strings.Split(payload, " ")

	var mentions []string
	for _, word := range input {
		if strings.HasPrefix(word, "@") {
			mentions = append(mentions, word)
		}
	}

	return mentions
}

func ResolveUserGroups(mentions []string, api *slack.Client) ([]string, error) {
	groups, err := api.GetUserGroups()
	if err != nil {
		return nil, err
	}

	var resolved []string
	var users []string
	for _, group := range groups {
		for _, mention := range mentions {
			if "@"+group.Handle == mention {
				resolved = append(resolved, mention)
				members, err := fetchUsersInGroup(api, group)
				if err != nil {
					return nil, err
				}

				users = append(users, members...)
				break
			}
		}
	}

	for _, m := range mentions {
		for _, r := range resolved {
			if m == r {
				break
			}
		}
		users = append(users, m)
	}

	return users, nil
}

func fetchUsersInGroup(api *slack.Client, group slack.UserGroup) ([]string, error) {
	members, err := api.GetUserGroupMembers(group.ID)
	if err != nil {
		return nil, err
	}

	return members, nil
}
