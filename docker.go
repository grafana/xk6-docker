package docker

import (
	"github.com/loadimpact/k6/js/modules"
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/docker", &Docker{
		Version: version,
	})

	containers := Containers{Version: version}
	containers.SetupClient()
	modules.Register("k6/x/docker/containers", &containers)
}

// Docker is the main export of k6 docker extension
type Docker struct {
	Version string
}
