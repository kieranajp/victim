package handler

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func ExtractUsers(text string) []string {
	r := regexp.MustCompile(`<([^<|>]*)[\|>]`)
	m := r.FindAllStringSubmatch(text, -1)

	var users []string
	for _, v := range m {
		users = append(users, v[1])
	}
	return users
}

func PickRandomUser(users []string) string {
	log.Info().
		Str("Users", strings.Join(users, ",")).
		Msg("Picking random user")

	rand.Seed(time.Now().Unix())
	user := users[rand.Intn(len(users))]
	return user
}
