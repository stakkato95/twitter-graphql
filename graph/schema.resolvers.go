package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/stakkato95/twitter-service-graphql/graph/generated"
	"github.com/stakkato95/twitter-service-graphql/graph/model"
	"github.com/stakkato95/twitter-service-graphql/http/middleware"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	return r.UserService.Create(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return r.UserService.Authenticate(input)
}

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*model.Tweet, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid authorization")
	}

	return r.TweetService.CreateTweet(input, int(user.Id))
}

func (r *mutationResolver) Subscribe(ctx context.Context, input model.NewSubscription) (string, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return "", errors.New("invalid authorization")
	}

	return r.TweetService.Subscribe(int(user.Id), input.To)
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid authorization")
	}

	return r.TweetService.GetTweets(int(user.Id))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
