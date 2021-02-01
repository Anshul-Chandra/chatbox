package v1

import (
	"context"

	"github.com/chatbox/proto/gen/v1/user"
	"go.uber.org/zap"
)

// UserService is the grpc service for user
type UserService struct {
}

// Get fetches user details for a given user
func (svc *UserService) Get(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	sugaredLogger.Infof("Received new request: %+v", req)
	resp := &user.GetUserResponse{
		User: &user.User{
			Id:        req.GetUserId(),
			FirstName: "First",
			LastName:  "User",
			UserEmail: "first-user@gmail.com",
		},
	}
	return resp, nil
}

// List lists all the users associated with a given user or group
func (svc *UserService) List(ctx context.Context, req *user.ListUserRequest) (*user.ListUserResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	sugaredLogger.Infof("Received new request: %+v", req)
	resp := &user.ListUserResponse{
		Users: []*user.User{
			{
				Id:        "1",
				FirstName: "First",
				LastName:  "User",
				UserEmail: "first-user@gmail.com",
			},
			{
				Id:        "2",
				FirstName: "Second",
				LastName:  "User",
				UserEmail: "second-user@gmail.com",
			},
		},
	}

	return resp, nil
}
