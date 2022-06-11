package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/stakkato95/twitter-service-graphql/http/service"
)

type Resolver struct {
	UserService  service.UserService
	TweetService service.TweetService
}
