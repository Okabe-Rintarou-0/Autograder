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
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	ChangePassword(ctx context.Context, userID uint, newPassword string) error
	UpdateCompilationInfo(ctx context.Context, userID uint, request *request.UpdateCompilationInfoRequest) error
	ListUsers(ctx context.Context, keyword string, page *entity.Page) (*response.ListUsersResponse, error)
	Register(ctx context.Context, request *request.RegisterRequest) (*response.RegisterResponse, error)
	ImportCanvasUsers(ctx context.Context, courseID int64) (*response.ImportCanvasUsers, error)
}
