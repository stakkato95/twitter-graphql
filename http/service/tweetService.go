package service

import (
	"github.com/stakkato95/twitter-service-graphql/graph/model"
	"github.com/stakkato95/twitter-service-graphql/http/domain"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

type TweetService interface {
	CreateTweet(model.NewTweet, int) (*model.Tweet, error)
	GetTweets(int) ([]*model.Tweet, error)
	Subscribe(int, int) (string, error)
}

type defaultTweetService struct {
	repo domain.TweetRepo
}

func NewTweetService(repo domain.TweetRepo) TweetService {
	return &defaultTweetService{repo}
}

func (s *defaultTweetService) CreateTweet(tweet model.NewTweet, userId int) (*model.Tweet, error) {
	tweetDto := dto.TweetToDto(tweet, userId)
	createdTweet, err := s.repo.CreateTweet(tweetDto)
	if err != nil {
		return nil, err
	}

	return dto.TweetDtoToGraphql(*createdTweet), nil
}

func (s *defaultTweetService) GetTweets(userId int) ([]*model.Tweet, error) {
	tweetsDto, err := s.repo.GetTweets(userId)
	if err != nil {
		return nil, err
	}

	tweets := make([]*model.Tweet, len(tweetsDto))
	for i, tweetDto := range tweetsDto {
		tweets[i] = dto.TweetDtoToGraphql(tweetDto)
	}

	return tweets, nil
}

func (s *defaultTweetService) Subscribe(from int, to int) (string, error) {
	result, err := s.repo.Subscribe(dto.SubscriptionDto{From: from, To: to})
	if err != nil {
		return "", err
	}
	return result, nil
}
