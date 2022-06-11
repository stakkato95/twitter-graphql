package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

var tweetsService = "http://" + config.AppConfig.TweetsService

type TweetRepo interface {
	CreateTweet(*dto.Tweet) (*dto.Tweet, error)
	GetTweets(int) ([]dto.Tweet, error)
	Subscribe(dto.SubscriptionDto) (string, error)
}

type defaultTweetRepo struct {
}

func NewTweetRepo() TweetRepo {
	return &defaultTweetRepo{}
}

func (r *defaultTweetRepo) CreateTweet(tweet *dto.Tweet) (*dto.Tweet, error) {
	jsonData, err := json.Marshal(tweet)
	if err != nil {
		logger.Fatal("can not encode tweet: " + err.Error())
		return nil, err
	}

	response, err := http.DefaultClient.Post(tweetsService+"/tweets", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("POST request to tweets service failed: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseDto := dto.ResponseDto{}
	if err := json.NewDecoder(response.Body).Decode(&responseDto); err != nil {
		logger.Fatal("can not decode response from tweets service: " + err.Error())
		return nil, err
	}

	if responseDto.Error != "" {
		return nil, errors.New(responseDto.Error)
	}

	jsonData, err = json.Marshal(responseDto.Data)
	if err != nil {
		return nil, errors.New("can not marshal tweet data: " + err.Error())
	}

	createdTweet := dto.Tweet{}
	if err := json.NewDecoder(bytes.NewBuffer(jsonData)).Decode(&createdTweet); err != nil {
		logger.Fatal("can not decode tweet from data: " + err.Error())
		return nil, err
	}

	return &createdTweet, nil
}

func (r *defaultTweetRepo) GetTweets(userId int) ([]dto.Tweet, error) {
	response, err := http.DefaultClient.Get(fmt.Sprintf(tweetsService+"/tweets/%d", userId))
	if err != nil {
		logger.Fatal("GET request to tweets service failed: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseDto := dto.ResponseDto{}
	if err := json.NewDecoder(response.Body).Decode(&responseDto); err != nil {
		logger.Fatal("can not decode response from tweets service: " + err.Error())
		return nil, err
	}

	if responseDto.Error != "" {
		return nil, errors.New(responseDto.Error)
	}

	jsonData, err := json.Marshal(responseDto.Data)
	if err != nil {
		return nil, errors.New("can not marshal tweets data: " + err.Error())
	}

	tweets := []dto.Tweet{}
	if err := json.NewDecoder(bytes.NewBuffer(jsonData)).Decode(&tweets); err != nil {
		logger.Fatal("can not decode tweets from data: " + err.Error())
		return nil, err
	}

	return tweets, nil
}

func (r *defaultTweetRepo) Subscribe(subscription dto.SubscriptionDto) (string, error) {
	jsonData, err := json.Marshal(subscription)
	if err != nil {
		logger.Fatal("can not encode subscription: " + err.Error())
		return "", err
	}

	response, err := http.DefaultClient.Post(tweetsService+"/subscription", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("POST request to tweets service failed: " + err.Error())
		return "", err
	}
	defer response.Body.Close()

	responseDto := dto.ResponseDto{}
	if err := json.NewDecoder(response.Body).Decode(&responseDto); err != nil {
		logger.Fatal("can not decode response from tweets service: " + err.Error())
		return "", err
	}

	if responseDto.Error != "" {
		return "", errors.New(responseDto.Error)
	}

	result, ok := responseDto.Data.(string)
	if !ok {
		logger.Fatal("can not decode string from data")
		return "", errors.New("can not decode string from data")
	}

	return result, nil
}
