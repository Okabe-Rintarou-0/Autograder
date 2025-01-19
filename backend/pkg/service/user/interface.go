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
}
