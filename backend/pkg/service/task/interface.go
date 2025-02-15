package task

import (
	"autograder/pkg/model/request"
	"context"
	"io"

	"autograder/pkg/model/entity"
	"autograder/pkg/model/response"
)

type Service interface {
	RunApp(ctx context.Context, info *entity.AppInfo) error
	SubmitApp(ctx context.Context, info *entity.AppInfo) (entity.SubmitAppResult, error)
	ListAppTasks(ctx context.Context, user *entity.User, req *request.ListAppRunTasksRequest) (*response.ListAppTasksResponse, error)
	GetLogFile(ctx context.Context, uuid, logType string) (io.ReadCloser, error)
}
