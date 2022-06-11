package service

import (
	"fmt"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/graph/model"
	"github.com/stakkato95/twitter-service-graphql/http/domain"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

type UserService interface {
	Create(model.NewUser) (string, error)
	Authenticate(model.Login) (string, error)
	Authorize(string) (*dto.UserDto, error)
}

type defaultUserService struct {
	repo domain.UserRepo
}

func NewUserService(repo domain.UserRepo) UserService {
	return &defaultUserService{repo}
}

func (s *defaultUserService) Create(user model.NewUser) (string, error) {
	userDto := dto.ToDtoFromUser(user)
	newUserDto, err := s.repo.Create(&userDto)
	if err != nil {
		return "", err
	}

	logger.Info(fmt.Sprintf("created user: %#v", newUserDto))
	return newUserDto.Token.Token, nil
}

func (s *defaultUserService) Authenticate(login model.Login) (string, error) {
	userDto := dto.ToDtoFromLogin(login)
	tokenDto, err := s.repo.Authenticate(&userDto)
	if err != nil {
		return "", err
	}

	return tokenDto.Token, nil
}

func (s *defaultUserService) Authorize(token string) (*dto.UserDto, error) {
	user, err := s.repo.Authorize(&dto.TokenDto{Token: token})
	if err != nil {
		return nil, err
	}

	return user, nil
	// return &dto.UserDto{
	// 	Id:       1,
	// 	Username: "user1",
	// 	Password: "pass",
	// }, nil
}
