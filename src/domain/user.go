package domain

import (
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
}

func NewUser(name string) User {
	u := User{}
	u.ID = u.generateID()
	u.Name = name
	return u
}

func (u User) generateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000)
}
