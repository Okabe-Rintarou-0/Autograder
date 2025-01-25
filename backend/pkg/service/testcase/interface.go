package testcase

import (
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"context"
)

type Service interface {
	Sync(ctx context.Context) error
	ListTestcases(ctx context.Context) ([]*response.Testcase, error)
	BatchUpdateTestcases(ctx context.Context, request *request.BatchUpdateTestcaseRequest) (*response.BatchUpdateTestcaseResponse, error)
}
