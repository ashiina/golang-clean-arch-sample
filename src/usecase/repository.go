package usecase

import "../domain"

type TweetRepository interface {
	Save(domain.Tweet) error
	GetAll() (domain.Tweets, error)
}

type UserRepository interface {
	Save(domain.User) error
	Get(userID int) (domain.User, error)
}
