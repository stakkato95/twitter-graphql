package domain

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

var usersService = "http://" + config.AppConfig.UsersService

type httpUserRepo struct {
}

func NewHttpUserRepo() UserRepo {
	return &httpUserRepo{}
}

func (r *httpUserRepo) Create(user *dto.UserDto) (*dto.NewUserDto, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		logger.Fatal("can not encode user: " + err.Error())
		return nil, err
	}

	response, err := http.DefaultClient.Post(usersService+"/debug/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("POST request to users service failed: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	newUser := dto.NewUserDto{}
	if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
		logger.Fatal("can not decode response from user service: " + err.Error())
		return nil, err
	}

	return &newUser, nil
}

func (r *httpUserRepo) Authenticate(user *dto.UserDto) (*dto.TokenDto, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		logger.Fatal("can not encode user: " + err.Error())
		return nil, err
	}

	response, err := http.DefaultClient.Post(usersService+"/debug/auth", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("POST request to users service failed: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	tokenDto := dto.TokenDto{}
	if err := json.NewDecoder(response.Body).Decode(&tokenDto); err != nil {
		logger.Fatal("can not decode response from user service: " + err.Error())
		return nil, err
	}

	return &tokenDto, nil
}

func (r *httpUserRepo) Authorize(*dto.TokenDto) (*dto.UserDto, error) {
	// stub
	stubUser := &dto.UserDto{
		Id:       1,
		Username: "userStub",
	}
	return stubUser, nil
}
