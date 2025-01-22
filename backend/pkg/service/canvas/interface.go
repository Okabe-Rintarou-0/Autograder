package canvas

import (
	"context"

	"autograder/pkg/model/entity/canvas"
)

type Service interface {
	ListCourses(ctx context.Context) ([]*canvas.Course, error)
}
