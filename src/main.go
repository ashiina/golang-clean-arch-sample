package main

import (
	"./domain"
	"./infra"
	"./usecase"
	"fmt"
)

func main() {
	fmt.Println("Running")

	u := domain.NewUser("Uncle Bob")

	tr := infra.NewTweetRepository()

	ptu := usecase.PostTweetUsecase{
		TweetRepository: tr,
	}
	text := "This is my first tweet!"
	err := ptu.Post(u, text)
	if err != nil {
		panic("Couldn't post tweet...")
	}

	gtu := usecase.GetTweetUsecase{
		TweetRepository: tr,
	}
	tweets := gtu.GetTweetsForUser(&u)

	for _, t := range tweets {
		fmt.Println(t.Text)
	}
}
