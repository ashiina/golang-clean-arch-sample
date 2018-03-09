package usecase

import "../domain"

type GetTweetUsecase struct {
	TweetRepository TweetRepository
}

func (ptu *GetTweetUsecase) GetTweetsForUser(u *domain.User) domain.Tweets {
	tweets := ptu.TweetRepository.GetTweetsForUser(u.ID)

	// some logic here...

	return tweets
}
