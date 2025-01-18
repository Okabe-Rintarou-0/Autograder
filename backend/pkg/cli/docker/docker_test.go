package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	c "github.com/smartystreets/goconvey/convey"

	"autograder/pkg/entity"
)

const (
	TestImageName = "maven:3.8.4-openjdk-17"
)

func TestDockerPullImage(t *testing.T) {
	c.Convey("Pull Image", t, func() {
		client := NewClient()
		ctx := context.Background()
		err := client.PullImage(ctx, "nginx:latest")
		c.So(err, c.ShouldBeNil)
	})
}

func TestDockerRunContainer(t *testing.T) {
	c.Convey("Run Container", t, func() {
		client := NewClient()
		ctx := context.Background()
		_ = client.RemoveContainer(ctx, "test_docker")
		id, err := client.RunContainer(ctx, &entity.DockerCreateConfig{
			ImageName:      TestImageName,
			ContainerName:  "test_docker",
			PortBindings:   map[string]string{"8080": "8080"},
			Commands:       []string{"/bin/bash"},
			VolumeBindings: map[string]string{"/Users/lucas/bookstore-backend/bookstore-backend": "/app"},
		})
		if err != nil {
			defer func() {
				_ = client.RemoveContainer(ctx, id)
			}()
		}
		c.So(err, c.ShouldBeNil)
		c.So(len(id), c.ShouldBeGreaterThan, 0)
	})
}

func TestDockerExecuteContainer(t *testing.T) {
	c.Convey("Execute Container", t, func() {
		client := NewClient()
		ctx := context.Background()
		commands := entity.NewBashCommandsBuilder().
			NewCommand("cd", "/app").
			NewCommand("mvn", "clean", "package").
			// NewCommand("java", "-jar", "target/*.jar").
			Build()
		fmt.Println(commands)
		reader, _ := client.ExecuteContainer(ctx, "test_docker", commands)
		io.Copy(os.Stdout, reader)
	})
}
