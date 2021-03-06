package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/loadimpact/k6/js/modules"
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/docker", &Docker{
		Version: version,
	})
}

// Docker is the main export of k6 docker extension
type Docker struct {
	Version string
}

// ListContainers works as Docker ps command
func (d *Docker) ListContainers(options types.ContainerListOptions) ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return []types.Container{}, err
	}

	containers, err := cli.ContainerList(context.Background(), options)
	return containers, err
}

// StartContainer works as Docker start command
func (d *Docker) StartContainer(containerID string, options types.ContainerStartOptions) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return err
	}

	return cli.ContainerStart(context.Background(), containerID, options)
}

// StopContainer works as Docker stop command
func (d *Docker) StopContainer(containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return err
	}

	// TODO: Add timeout option support
	timeout := 0 * time.Second

	return cli.ContainerStop(context.Background(), containerID, &timeout)
}
