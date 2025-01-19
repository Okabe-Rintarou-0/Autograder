package hurl

import (
	"context"
	"fmt"
	"os/exec"

	"autograder/pkg/config"

	"github.com/sirupsen/logrus"
)

type daoImpl struct{}

func NewDAO() *daoImpl {
	return &daoImpl{}
}

func (d *daoImpl) RunAllTests(ctx context.Context) error {
	cmd := exec.Command("hurl", "--test", config.Instance.TestcasesDir)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests]: error %+v, output: %s", err, string(output))
		return err
	}
	fmt.Println(string(output))
	return nil
}
