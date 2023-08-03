package docker

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Containers is a structure with all docker containers functions
type Containers struct {
	Client *client.Client
}

// SetupClient for filling in Docker client instance
func (d *Containers) SetupClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	d.Client = cli
}

// List works as Docker ps command
func (d *Containers) List(options types.ContainerListOptions) ([]types.Container, error) {
	containers, err := d.Client.ContainerList(context.Background(), options)
	return containers, err
}

// Start works as Docker start command
func (d *Containers) Start(containerID string) error {
	return d.Client.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

// Stop works as Docker stop command
func (d *Containers) Stop(containerID string) error {
	// TODO: Add timeout option support
	timeout := 0
	options := container.StopOptions{
		Timeout: &timeout,
	}

	return d.Client.ContainerStop(context.Background(), containerID, options)
}

// Pause works as Docker pause command
func (d *Containers) Pause(containerID string) error {
	return d.Client.ContainerPause(context.Background(), containerID)
}

// Unpause works as Docker unpause command
func (d *Containers) Unpause(containerID string) error {
	return d.Client.ContainerUnpause(context.Background(), containerID)
}

// Logs works as Docker logs command
func (d *Containers) Logs(containerID string, options types.ContainerLogsOptions) (string, error) {
	response, err := d.Client.ContainerLogs(context.Background(), containerID, options)
	buf := new(bytes.Buffer)

	if err != nil {
		return "", err
	}

	_, err = buf.ReadFrom(response)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Inspect works as Docker inspect command
func (d *Containers) Inspect(containerID string) (types.ContainerJSON, error) {
	return d.Client.ContainerInspect(context.Background(), containerID)
}

// Exec works as Docker exec command
func (d *Containers) Exec(containerID string, config types.ExecConfig) error {
	response, err := d.Client.ContainerExecCreate(context.Background(), containerID, config)
	if err != nil {
		return err
	}
	return d.Client.ContainerExecStart(context.Background(), response.ID, types.ExecStartCheck{})
}
