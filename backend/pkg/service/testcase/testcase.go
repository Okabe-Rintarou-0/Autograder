package testcase

import (
	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/messages"
	"autograder/pkg/model/assembler"
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type ServiceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *ServiceImpl {
	return &ServiceImpl{groupDAO}
}

func (s *ServiceImpl) BatchUpdateTestcases(ctx context.Context, request *request.BatchUpdateTestcaseRequest) (*response.BatchUpdateTestcaseResponse, error) {
	r := &response.BatchUpdateTestcaseResponse{}
	if len(request.Data) == 0 {
		r.BaseResp = response.NewSucceedBaseResp(messages.ModifySucceed)
		return r, nil
	}

	models := utils.Map(request.Data, assembler.ConvertTestcaseRequestToDBM)
	err := s.groupDAO.TestcaseDAO.Save(ctx, models...)
	if err != nil {
		return nil, err
	}

	r.BaseResp = response.NewSucceedBaseResp(messages.ModifySucceed)
	return r, nil
}

func (s *ServiceImpl) Sync(ctx context.Context) error {
	testcaseFiles, err := utils.GetAllFileNames(config.Instance.TestcasesDir, ".hurl")
	if err != nil {
		logrus.Errorf("[Testcase Service] get test case files err: %v", err)
		return err
	}

	testcaseModels, err := s.groupDAO.TestcaseDAO.FindAll(ctx, nil)
	if err != nil {
		logrus.Errorf("[Testcase Service] get testcase models err: %v", err)
		return err
	}

	testcaseFileMap := utils.IntoSet(testcaseFiles, func(v string) string {
		return v
	})

	nonExistentModels := utils.Filter(testcaseModels, func(v *dbm.Testcase) bool {
		_, ok := testcaseFileMap[v.Name]
		return !ok
	})

	nonExistentModelNames := utils.Map(nonExistentModels, func(v *dbm.Testcase) string {
		return v.Name
	})

	if len(nonExistentModelNames) > 0 {
		err = s.groupDAO.TestcaseDAO.DeleteAll(ctx, &dbm.TestcaseFilter{
			Names: nonExistentModelNames,
		})
	}
	if err != nil {
		logrus.Errorf("[Testcase Service] delete test case models err: %v", err)
		return err
	}

	saveModels := utils.Map(testcaseFiles, func(v string) *dbm.Testcase {
		bytes, _ := os.ReadFile(v)
		return &dbm.Testcase{
			Name:    v,
			Status:  dbm.Active,
			Content: string(bytes),
		}
	})
	err = s.groupDAO.TestcaseDAO.SaveIfNotExist(ctx, saveModels...)
	if err != nil {
		logrus.Errorf("[Testcase Service] save test case models err: %v", err)
		return err
	}
	return nil
}

func (s *ServiceImpl) ListTestcases(ctx context.Context) ([]*response.Testcase, error) {
	models, err := s.groupDAO.TestcaseDAO.FindAll(ctx, nil)
	if err != nil {
		logrus.Errorf("[Testcase Service] get testcase models err: %v", err)
		return nil, err
	}

	return utils.Map(models, assembler.ConvertTestcaseDbmToResponse), nil
}
