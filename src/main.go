package main

import (
	"./domain"
	"./usecase"
	"fmt"
)

func main() {
	fmt.Println("Running")

	u := domain.NewUser("Uncle Bob")

	text := "This is my first tweet!"
	ptu := usecase.PostTweetUsecase{}
	err := ptu.Post(u, text)
	if err != nil {
		panic("Couldn't post tweet...")
	}

}
