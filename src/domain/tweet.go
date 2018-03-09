package domain

import (
	"math/rand"
	"time"
)

type Tweets []Tweet

type Tweet struct {
	ID     int
	UserID int
	Text   string
}

func NewTweet(userID int, text string) Tweet {
	t := Tweet{}
	t.ID = t.generateID()
	t.UserID = userID
	t.Text = text
	if !t.isValidLength(text) {
		panic("Invalid tweet length") // FIXME
	}

	return t
}

func (t Tweet) generateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000)
}

func (t Tweet) isValidLength(text string) bool {
	return (len(text) <= 140)
}
