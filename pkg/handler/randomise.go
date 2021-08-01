package handler

import (
	"math/rand"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func PickRandomUser(users string) string {
	log.Info().
		Str("Users", users).
		Msg("Picking random user")

	u := strings.Split(users, ",")

	rand.Seed(time.Now().Unix())
	user := u[rand.Intn(len(u))]
	return user
}
