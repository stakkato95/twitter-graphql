package domain

import (
	"context"
	"time"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/stakkato95/twitter-service-graphql/proto"
)

const timeout = 10

type grpcUserRepo struct {
	client pb.UsersServiceClient
}

func NewGrpcUserRepo() UserRepo {
	conn, err := grpc.Dial(config.UsersGrpc(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("can not listen to users grpc server: " + err.Error())
	}
	return &grpcUserRepo{pb.NewUsersServiceClient(conn)}
}

func (r *grpcUserRepo) Create(user *dto.UserDto) (*dto.NewUserDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	newUser, err := r.client.CreateUser(ctx, dto.NewUserToProto(user))
	if err != nil {
		logger.Fatal("can not create user via users grpc interface: " + err.Error())
	}
	return dto.NewUserToDto(newUser), nil
}

func (r *grpcUserRepo) Authenticate(user *dto.UserDto) (*dto.TokenDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	token, err := r.client.AuthUser(ctx, dto.NewUserToProto(user))
	if err != nil {
		logger.Fatal("can not authenticate user via users grpc interface: " + err.Error())
	}
	return dto.TokenToDto(token), nil
}

func (r *grpcUserRepo) Authorize(token *dto.TokenDto) (*dto.UserDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	user, err := r.client.AuthUserByToken(ctx, &pb.Token{Token: token.Token})
	if err != nil {
		return nil, err
	}
	return dto.UserToDto(user), nil
}
