package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/client"
	"github.com/sirupsen/logrus"

	"autograder/pkg/entity"
)

type clientImpl struct {
	innerCli *client.Client
}

func NewClient() *clientImpl {
	innerCli, err := client.NewClientWithOpts()
	if err != nil {
		logrus.Fatalf("[Docker client][NewClient] inner client create error %+v", err)
		panic(err)
	}
	return &clientImpl{innerCli}
}

func (c *clientImpl) PullImage(ctx context.Context, imageName string) error {
	options := image.PullOptions{}
	resp, err := c.innerCli.ImagePull(ctx, imageName, options)
	if err != nil {
		logrus.Errorf("[Docker client][PullImage] inner client call ImagePull error %+v", err)
		return err
	}
	defer resp.Close()
	if _, err := io.ReadAll(resp); err != nil {
		logrus.Errorf("[Docker client][PullImage] io.ReadAll error %+v", err)
		return err
	}
	return nil
}

func (c *clientImpl) RemoveContainer(ctx context.Context, containerID string) error {
	err := c.innerCli.ContainerRemove(ctx, containerID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		logrus.Errorf("[Docker client][RemoveContainer] inner client call ContainerRemove error %+v", err)
	}
	return err
}

func (c *clientImpl) RunContainer(ctx context.Context, config *entity.DockerCreateConfig) (string, error) {
	exposedPorts := nat.PortSet{}
	portBindings := nat.PortMap{}
	for containerPort, hostPort := range config.PortBindings {
		port, err := nat.NewPort("tcp", containerPort)
		if err != nil {
			return "", err
		}
		exposedPorts[port] = struct{}{}
		portBindings[port] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		}
	}
	containerConfig := &container.Config{
		Image:        config.ImageName,
		ExposedPorts: exposedPorts,
		Tty:          true,
		Cmd:          config.Commands,
	}
	var mounts []mount.Mount
	for source, target := range config.VolumeBindings {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: source,
			Target: target,
		})
	}
	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
		Mounts:       mounts,
	}
	networkConfig := &network.NetworkingConfig{}

	resp, err := c.innerCli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, config.ContainerName)
	if err != nil {
		return "", err
	}

	if err := c.innerCli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	logrus.Infof("[Docker client][RunContainer] run container %s succeeded, config: %s", resp.ID, config.FormatString())
	return resp.ID, nil
}

func (c *clientImpl) ExecuteContainer(ctx context.Context, containerId string, commands []string) error {
	execConfig := container.ExecOptions{
		Cmd:          commands,
		AttachStdout: true,
		AttachStderr: false,
		AttachStdin:  false,
		Tty:          true,
	}
	execResp, err := c.innerCli.ContainerExecCreate(ctx, containerId, execConfig)
	if err != nil {
		logrus.Errorf("[Docker client][ExecuteContainer] call innerCli.ContainerExecCreate error %+v", err)
		return err
	}
	attachResp, err := c.innerCli.ContainerExecAttach(context.Background(), execResp.ID, container.ExecStartOptions{
		Detach: false,
	})
	if err != nil {
		logrus.Errorf("[Docker client][ExecuteContainer] call innerCli.ContainerExecAttach error %+v", err)
		return err
	}
	defer attachResp.Close()
	_, err = io.Copy(os.Stdout, attachResp.Reader)
	if err != nil {
		logrus.Errorf("[Docker client][ExecuteContainer] call io.Copy error %+v", err)
		return err
	}
	return nil
}
