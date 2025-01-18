package hurl

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
)

type daoImpl struct{}

func NewDAO() *daoImpl {
	return &daoImpl{}
}

func (d *daoImpl) RunAllTests(ctx context.Context) error {
	cmd := exec.Command("hurl", "--test", config.TestcasesDir)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("[Hurl DAO][RunAllTests]: error %+v, output: %s", err, string(output))
		return err
	}
	fmt.Println(string(output))
	return nil
}
