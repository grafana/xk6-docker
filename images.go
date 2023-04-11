package docker

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Images is a structure with all docker Image functions
type Images struct {
	Client *client.Client
}

// SetupClient for filling in Docker client instance
func (d *Images) SetupClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	d.Client = cli
}

// List works as Docker image ls command
func (d *Images) List(options types.ImageListOptions) ([]types.ImageSummary, error) {
	containers, err := d.Client.ImageList(context.Background(), options)
	return containers, err
}

// Pull works as Docker pull
func (d *Images) Pull(refStr string, options types.ImagePullOptions) (string, error) {
	response, err := d.Client.ImagePull(context.Background(), refStr, options)
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

// Remove works as Docker image rm
func (d *Images) Remove(imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	return d.Client.ImageRemove(context.Background(), imageID, options)
}

// Inspect works as Docker image inspect
func (d *Images) Inspect(imageID string) (types.ImageInspect, error) {
	result, _, err := d.Client.ImageInspectWithRaw(context.Background(), imageID)
	return result, err
}
