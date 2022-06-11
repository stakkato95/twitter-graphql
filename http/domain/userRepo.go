package domain

import (
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

type UserRepo interface {
	Create(*dto.UserDto) (*dto.NewUserDto, error)
	Authenticate(*dto.UserDto) (*dto.TokenDto, error)
	Authorize(*dto.TokenDto) (*dto.UserDto, error)
}
