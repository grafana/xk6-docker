package docker

import (
	"bytes"
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Containers is a structure with all docker containers functions
type Containers struct {
	Version string
	Client  *client.Client
}

// SetupClient for filling in Docker client instance
func (d *Containers) SetupClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	d.Client = cli
}

// ListContainers works as Docker ps command
func (d *Containers) ListContainers(options types.ContainerListOptions) ([]types.Container, error) {
	containers, err := d.Client.ContainerList(context.Background(), options)
	return containers, err
}

// StartContainer works as Docker start command
func (d *Containers) StartContainer(containerID string, options types.ContainerStartOptions) error {
	return d.Client.ContainerStart(context.Background(), containerID, options)
}

// StopContainer works as Docker stop command
func (d *Containers) StopContainer(containerID string) error {
	// TODO: Add timeout option support
	timeout := 0 * time.Second

	return d.Client.ContainerStop(context.Background(), containerID, &timeout)
}

// PauseContainer works as Docker pause command
func (d *Containers) PauseContainer(containerID string) error {
	return d.Client.ContainerPause(context.Background(), containerID)
}

// UnpauseContainer works as Docker unpause command
func (d *Containers) UnpauseContainer(containerID string) error {
	return d.Client.ContainerUnpause(context.Background(), containerID)
}

// Logs works as Docker logs command
func (d *Containers) Logs(containerID string, options types.ContainerLogsOptions) (string, error) {
	response, err := d.Client.ContainerLogs(context.Background(), containerID, options)
	buf := new(bytes.Buffer)

	if err != nil {
		return "", err
	}

	buf.ReadFrom(response)
	return buf.String(), nil
}

// InspectContainer works as Docker inspect command
func (d *Containers) InspectContainer(containerID string) (types.ContainerJSON, error) {
	return d.Client.ContainerInspect(context.Background(), containerID)
}
