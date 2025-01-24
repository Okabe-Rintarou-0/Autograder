package canvas

import (
	"context"

	"autograder/pkg/dao"
	"autograder/pkg/model/entity/canvas"
)

type ServiceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *ServiceImpl {
	return &ServiceImpl{groupDAO}
}

func (s *ServiceImpl) ListCourses(ctx context.Context) ([]*canvas.Course, error) {
	return s.groupDAO.CanvasDAO.ListCourses(ctx)
}

func (s *ServiceImpl) ListAssignments(ctx context.Context, courseID int64) ([]*canvas.Assignment, error) {
	return s.groupDAO.CanvasDAO.ListAssignments(ctx, courseID)
}

func (s *ServiceImpl) ListAssignmentSubmissions(ctx context.Context, courseID, assignmentID int64) ([]*canvas.Submission, error) {
	return s.groupDAO.CanvasDAO.ListAssignmentSubmissions(ctx, courseID, assignmentID)
}

func (s *ServiceImpl) ListCourseUsers(ctx context.Context, courseID int64) ([]*canvas.User, error) {
	return s.groupDAO.CanvasDAO.ListCourseUsers(ctx, courseID)
}
