package handler

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// ExtractUsers takes in the incoming text from a Slack command, and finds all the linkable @user entities within.
func ExtractUsers(text string) []string {
	r := regexp.MustCompile(`<([^<|>]*)[\|>]`)
	m := r.FindAllStringSubmatch(text, -1)

	var users []string
	for _, v := range m {
		users = append(users, fmt.Sprintf("<%s>", v[1]))
	}
	return users
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
