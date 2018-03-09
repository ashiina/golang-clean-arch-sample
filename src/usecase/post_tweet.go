package usecase

import "../domain"

type PostTweetUsecase struct {
	TweetRepository TweetRepository
}

func (ptu *PostTweetUsecase) Post(u domain.User, text string) error {
	t := domain.NewTweet(u.ID, text)
	err := ptu.TweetRepository.Save(t)
	if err != nil {
		return err
	}
	return nil
}
