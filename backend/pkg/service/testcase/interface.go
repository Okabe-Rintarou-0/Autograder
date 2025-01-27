package testcase

import (
	"context"

	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
)

type Service interface {
	Sync(ctx context.Context) error
	ListTestcases(ctx context.Context) ([]*response.Testcase, error)
	BatchUpdateTestcases(ctx context.Context, request *request.BatchUpdateTestcaseRequest) (*response.BatchUpdateTestcaseResponse, error)
}
