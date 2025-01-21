package user

import (
	"context"

	"autograder/pkg/model/entity"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
)

type Service interface {
	Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error)
	GetUser(ctx context.Context, userID uint) (*entity.User, error)
	ChangePassword(ctx context.Context, userID uint, newPassword string) error
	ListUsers(ctx context.Context, page *entity.Page) (*response.ListUsersResponse, error)
	Register(ctx context.Context, request *request.RegisterRequest) (*response.RegisterResponse, error)
}
