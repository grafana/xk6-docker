package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/loadimpact/k6/js/modules"
)

const version = "v0.0.2"

func init() {
	modules.Register("k6/x/docker", &Docker{
		Version: version,
	})
}

// Docker is the main export of k6 docker extension
type Docker struct {
	Version string
}

// Ps lists Docker containers
func (d *Docker) Ps() []string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	result := make([]string, 0, len(containers))

	for _, container := range containers {
		result = append(result, fmt.Sprintf("%s %s %s", container.ID[:10], container.Image, container.Command))
	}

	return result
}
