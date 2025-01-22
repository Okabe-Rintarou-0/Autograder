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
