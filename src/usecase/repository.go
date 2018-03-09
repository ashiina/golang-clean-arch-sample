package usecase

import "../domain"

type TweetRepository interface {
	Save(domain.Tweet) error
	GetTweetsForUser(userID int) domain.Tweets
}

type UserRepository interface {
	Save(domain.User) error
	Get(userID int) (domain.User, error)
}
