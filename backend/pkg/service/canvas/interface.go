package canvas

import (
	"context"

	"autograder/pkg/model/entity/canvas"
)

type Service interface {
	ListCourses(ctx context.Context) ([]*canvas.Course, error)
	ListAssignments(ctx context.Context, courseID int64) ([]*canvas.Assignment, error)
	ListAssignmentSubmissions(ctx context.Context, courseID, assignmentID int64) ([]*canvas.Submission, error)
	ListCourseUsers(ctx context.Context, courseID int64) ([]*canvas.User, error)
}
