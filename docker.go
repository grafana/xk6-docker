// Package docker provides a module implementation for manipulating docker resources using Javascript
package docker

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/docker", &Docker{})

	containers := Containers{}
	containers.SetupClient()
	modules.Register("k6/x/docker/containers", &containers)

	volumes := Volumes{}
	volumes.SetupClient()
	modules.Register("k6/x/docker/volumes", &volumes)

	networks := Networks{}
	networks.SetupClient()
	modules.Register("k6/x/docker/networks", &networks)

	images := Images{}
	images.SetupClient()
	modules.Register("k6/x/docker/images", &images)
}

// Docker is the main export of k6 docker extension
type Docker struct{}
