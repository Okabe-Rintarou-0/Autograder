package canvas

import (
	"context"

	"autograder/pkg/model/entity/canvas"
)

type DAO interface {
	ListCourses(ctx context.Context) ([]*canvas.Course, error)
	ListCourseUsers(ctx context.Context, courseID int64) ([]*canvas.User, error)
}
