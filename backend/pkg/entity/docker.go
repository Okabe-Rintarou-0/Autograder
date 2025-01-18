package entity

import "autograder/pkg/utils"

type DockerCreateConfig struct {
	ImageName      string
	ContainerName  string
	PortBindings   map[string]string
	Commands       []string
	VolumeBindings map[string]string
}

func (c *DockerCreateConfig) FormatString() string {
	return utils.FormatJsonString(c)
}
