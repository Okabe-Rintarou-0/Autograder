package file

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"autograder/pkg/model/constants"

	"github.com/sirupsen/logrus"

	"autograder/pkg/dal/cli/docker"
	"autograder/pkg/model/entity"
)

type DaoImpl struct {
	cli        docker.Client
	imageReady bool
}

func NewDao() *DaoImpl {
	return &DaoImpl{}
}

func (d *DaoImpl) Cleanup(ctx context.Context, info *entity.AppInfo) error {
	appDir := info.AppPath()
	return os.RemoveAll(appDir)
}

func (d *DaoImpl) Unzip(ctx context.Context, info *entity.AppInfo) error {
	r, err := zip.OpenReader(info.ZipFilePath())
	if err != nil {
		logrus.Errorf("[Unzip DAO][Unzip] call zip.OpenReader error %+v", err)
		return err
	}
	defer r.Close()

	appDir := info.AppPath()
	_ = os.RemoveAll(appDir)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		logrus.Errorf("[Unzip DAO][Unzip] call os.MkdirAll error %+v", err)
		return err
	}

	for _, f := range r.File {
		targetPath := filepath.Join(appDir, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, f.Mode()); err != nil {
				logrus.Errorf("[Unzip DAO][Unzip] call os.MkdirAll error %+v", err)
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			logrus.Errorf("[Unzip DAO][Unzip] call os.MkdirAll error %+v", err)
			return err
		}
		outFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DaoImpl) PrepareLogFile(ctx context.Context, info *entity.AppInfo) (io.Writer, io.Writer, error) {
	logDir := info.GetLogDir()
	if logDir == nil {
		return nil, nil, fmt.Errorf("nil log file")
	}
	stdout, err := logDir.GetWriter(constants.LogTypeStdout)
	if err != nil {
		return nil, nil, err
	}
	stderr, err := logDir.GetWriter(constants.LogTypeStderr)
	if err != nil {
		return nil, nil, err
	}
	return stdout, stderr, nil
}
