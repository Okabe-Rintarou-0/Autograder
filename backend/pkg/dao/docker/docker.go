package docker

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"

	"autograder/pkg/cli/docker"
	"autograder/pkg/entity"
)

var (
	jdkImageNameMap = map[int32]string{
		17: "maven:3.8.4-openjdk-17",
		11: "maven:3.8.4-openjdk-11",
	}
)

type daoImpl struct {
	cli        docker.Client
	imageReady bool
}

func NewDAO() *daoImpl {
	return &daoImpl{
		cli:        docker.NewClient(),
		imageReady: false,
	}
}

func (d *daoImpl) checkHTTP() bool {
	_, err := http.Get("http://localhost:8080")
	if err == nil {
		return true
	}
	return false
}

func (d *daoImpl) CompileAndRun(ctx context.Context, info *entity.AppInfo) (ContainerRemoveFn, error) {
	compileContainerImageName := jdkImageNameMap[info.JDKVersion]
	if !d.imageReady {
		err := d.cli.PullImage(ctx, compileContainerImageName)
		if err != nil {
			logrus.Errorf("[Docker DAO][RunCompileContainer] call client.PullImage error: %+v", err)
			return nil, err
		}
		d.imageReady = true
	}

	appPath := info.AppPath()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cachePath := filepath.Join(homeDir, "/.m2/repository")
	id, err := d.cli.RunContainer(ctx, &entity.DockerCreateConfig{
		ImageName:     compileContainerImageName,
		ContainerName: info.GetUUID(),
		PortBindings:  map[string]string{"8080": "8080"},
		Commands:      []string{"/bin/bash"},
		VolumeBindings: map[string]string{
			appPath:   "/app",
			cachePath: "/root/.m2/repository",
		},
	})
	if err != nil {
		logrus.Errorf("[Docker DAO][RunCompileContainer] call client.RunContainer error: %+v", err)
		return nil, err
	}

	doneCh := make(chan struct{})
	go func() {
		commands := entity.NewBashCommandsBuilder().
			NewCommand("cd", "/app").
			NewCommand("ls").
			NewCommand("mvn", "clean", "package").
			NewCommand("java", "-jar", "target/*.jar").
			Build()
		err := d.cli.ExecuteContainer(ctx, id, commands)
		if err != nil {
			logrus.Errorf("[Docker DAO][RunCompileContainer] call client.ExecuteContainer error: %+v", err)
		}
		doneCh <- struct{}{}
	}()

	ticker := time.NewTicker(time.Second)
	timeout := time.After(time.Minute)
	finished := false
	for !finished {
		select {
		case <-doneCh:
			logrus.Info("[Docker DAO][RunCompileContainer] ExecuteContainer finished")
			finished = true
		case <-ticker.C:
			if d.checkHTTP() {
				logrus.Info("[Docker DAO][RunCompileContainer] HTTP server is ready")
				finished = true
			} else {
				logrus.Info("[Docker DAO][RunCompileContainer] HTTP server is not ready")
			}
		case <-timeout:
			logrus.Info("[Docker DAO][RunCompileContainer] time out")
			finished = true
		}
	}

	return func() error {
		return d.cli.RemoveContainer(ctx, id)
	}, nil
}
