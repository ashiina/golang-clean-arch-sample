package infra

import "../domain"

type TweetRepository struct {
	tweets domain.Tweets
}

func NewTweetRepository() *TweetRepository {
	return &TweetRepository{}
}

func (tr *TweetRepository) Save(t domain.Tweet) error {
	tr.tweets = append(tr.tweets, t)
	return nil
}

func (tr *TweetRepository) GetTweetsForUser(userID int) domain.Tweets {

	// some filtering logic here...

	return tr.tweets
}
