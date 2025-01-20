package hurl

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/model/constants"
	"autograder/pkg/model/entity"
)

type DaoImpl struct{}

func NewDAO() *DaoImpl {
	return &DaoImpl{}
}

func (d *DaoImpl) RunAllTests(ctx context.Context, info *entity.AppInfo) (string, []*entity.HurlTestResult, error) {
	logDir := info.GetLogDir()
	reportDir := logDir.DirPath
	reportJsonPath := filepath.Join(reportDir, "report.json")

	err := os.Remove(reportJsonPath)
	if err != nil {
		logrus.Warnf("[Hurl DAO][RunAllTests] call os.ReadFile error %+v", err)
	}

	cmd := exec.Command("hurl", "--test", config.Instance.TestcasesDir, "--report-json", reportDir)

	writer, err := logDir.GetWriter(constants.LogTypeHurlTest)
	if err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests] call logDir.GetWriter error %+v", err)
		return "", nil, err
	}
	cmd.Stdout = writer
	cmd.Stderr = writer

	if err = cmd.Run(); err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests] run command error %+v", err)
		return "", nil, err
	}

	file, err := os.Open(reportJsonPath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
	}

	var results []*entity.HurlTestResult
	err = json.Unmarshal(bytes, &results)
	return string(bytes), results, err
}
