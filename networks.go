package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Networks is a structure with all docker Network functions
type Networks struct {
	Client *client.Client
}

// SetupClient for filling in Docker client instance
func (nw *Networks) SetupClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	nw.Client = cli
}

// List works as Docker network ls
func (nw *Networks) List(options types.NetworkListOptions) ([]types.NetworkResource, error) {
	return nw.Client.NetworkList(context.Background(), options)
}

// Create works as Docker network create
func (nw *Networks) Create(name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	return nw.Client.NetworkCreate(context.Background(), name, options)
}

// Remove works as Docker network remove
func (nw *Networks) Remove(networkID string) error {
	return nw.Client.NetworkRemove(context.Background(), networkID)
}
