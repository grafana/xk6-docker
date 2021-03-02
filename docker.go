package docker

import (
	"github.com/loadimpact/k6/js/modules"
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/docker", &Docker{
		Version: version,
	})
}

// Docker is the main export of k6 docker extension
type Docker struct {
	Version string
}
