package server

import (
	"context"

	"github.com/rezaindrag/mygrpcserver/model"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userService struct {
	model.UnimplementedUserServiceServer
}

// NewUserService initializes user gRPC service.
func NewUserService() model.UserServiceServer {
	return &userService{}
}

func (s userService) Fetch(context.Context, *emptypb.Empty) (*model.UserList, error) {
	users := &model.UserList{
		List: []*model.User{
			{
				Id:          "1",
				FullName:    "Reza Indra G",
				Email:       "rezaindra@gmail.com",
				Hobbies:     []string{"Reading", "Cycling"},
				CreatedTime: timestamppb.Now(),
			},
		},
	}

	return users, nil
}
