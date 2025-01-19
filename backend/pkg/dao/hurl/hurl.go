package hurl

import (
	"autograder/pkg/config"
	"autograder/pkg/model/constants"
	"autograder/pkg/model/entity"
	"autograder/pkg/utils"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

type daoImpl struct{}

func NewDAO() *daoImpl {
	return &daoImpl{}
}

func (d *daoImpl) RunAllTests(ctx context.Context, info *entity.AppInfo) ([]*entity.HurlTestResult, error) {
	logDir := info.GetLogDir()
	reportDir := logDir.DirPath
	reportJsonPath := filepath.Join(reportDir, "report.json")

	_, _ = os.ReadFile(reportJsonPath)

	cmd := exec.Command("hurl", "--test", config.Instance.TestcasesDir, "--report-json", reportDir)

	writer, err := logDir.GetWriter(constants.LogTypeHurlTest)
	if err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests] call logDir.GetWriter error %+v", err)
		return nil, err
	}
	cmd.Stdout = writer
	cmd.Stderr = writer

	if err = cmd.Run(); err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests] run command error %+v", err)
		return nil, err
	}

	file, err := os.Open(reportJsonPath)
	if err != nil {
		return nil, err
	}
	var results []*entity.HurlTestResult
	err = utils.UnmarshalReader(file, &results)
	return results, err
}
