package database

import (
	"errors"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/rs/zerolog/log"
)

func New() *bolt.DB {
	db, err := bolt.Open("./data/victim.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open database")
	}

	return db
}

func GetToken(db *bolt.DB, teamID string) (string, string, error) {
	var token string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tokens"))
		if b == nil {
			return nil
		}

		token = string(b.Get([]byte(teamID)))
		return nil
	})

	tokens := strings.Split(token, "|")
	if len(tokens) != 2 {
		return "", "", errors.New("token saved in invalid format")
	}

	return tokens[0], tokens[1], err
}

func SaveToken(db *bolt.DB, teamID, accessToken, botToken string) error {
	token := strings.Join([]string{accessToken, botToken}, "|")

	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tokens"))
		if err != nil {
			return err
		}

		return b.Put([]byte(teamID), []byte(token))
	})
}
