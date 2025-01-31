package docker

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"autograder/pkg/dal/cli/docker"
	"autograder/pkg/model/entity"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/sirupsen/logrus"
)

var (
	jdkImageNameMap = map[int32]string{
		17: "maven:3.8.4-openjdk-17",
		11: "maven:3.8.4-openjdk-11",
	}
)

type DaoImpl struct {
	cli        docker.Client
	imageReady bool
}

func NewDAO() *DaoImpl {
	return &DaoImpl{
		cli:        docker.NewClient(),
		imageReady: false,
	}
}

func (d *DaoImpl) checkHTTP() bool {
	_, err := http.Get("http://localhost:8080")
	if err == nil {
		return true
	}
	return false
}

func (d *DaoImpl) RemoveContainer(ctx context.Context, containerID string) error {
	return d.cli.RemoveContainer(ctx, containerID)
}

func (d *DaoImpl) CompileAndRun(ctx context.Context, info *entity.AppInfo, stdoutWriter, stderrWriter io.WriteCloser) (string, error) {
	compileContainerImageName := jdkImageNameMap[info.JDKVersion]
	if !d.imageReady {
		err := d.cli.PullImage(ctx, compileContainerImageName)
		if err != nil {
			logrus.Errorf("[Docker DAO][RunCompileContainer] call client.PullImage error: %+v", err)
			return "", err
		}
		d.imageReady = true
	}

	appPath := info.ProjectDirPath
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cachePath := filepath.Join(homeDir, "/.m2/repository")
	id, err := d.cli.RunContainer(ctx, &entity.DockerCreateConfig{
		ImageName:     compileContainerImageName,
		ContainerName: info.UUID,
		PortBindings:  map[string]string{"8080": "8080"},
		Commands:      []string{"/bin/bash"},
		VolumeBindings: map[string]string{
			appPath:   "/app",
			cachePath: "/root/.m2/repository",
		},
	})
	if err != nil {
		logrus.Errorf("[Docker DAO][RunCompileContainer] call client.RunContainer error: %+v", err)
		return "", err
	}

	doneCh := make(chan struct{})
	commands := entity.NewBashCommandsBuilder().
		NewCommand("cd", "/app").
		NewCommand("ls").
		NewCommand("mvn", "clean", "package").
		NewCommand("java", "-jar", "target/*.jar").
		Build()
	reader, err := d.cli.ExecuteContainer(ctx, id, commands)
	if err != nil {
		return id, err
	}
	go func() {
		defer func() {
			_ = stdoutWriter.Close()
			_ = stderrWriter.Close()
		}()
		// https://stackoverflow.com/questions/46478169/explain-and-remove-useless-bytes-at-the-start-of-docker-exec-response
		_, err := stdcopy.StdCopy(stdoutWriter, stderrWriter, reader)
		if err != nil {
			logrus.Errorf("[Docker DAO][RunCompileContainer] call io.Copy error: %+v", err)
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
	return id, nil
}
